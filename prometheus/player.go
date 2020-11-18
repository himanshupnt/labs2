package hangman

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	goodGuess = "😃"
	badGuess  = "😡"
	badEntry  = ' '
)

type Player struct {
	game    *Game
	address string
}

func NewPlayer(base string) *Player {
	return &Player{
		game:    &Game{},
		address: base,
	}
}

func (p *Player) Loop() error {
	if err := p.newGame(); err != nil {
		return err
	}

	turnsLeft, goodGuess := p.game.Tally.TurnsLeft, true
	for !p.gameOver() {
		print("\033[H\033[2J")
		fmt.Printf("\nYour Word: %10s\n", string(p.game.Tally.Letters))
		guess, err := p.prompt(goodGuess)
		if err != nil {
			fmt.Println("Bad entry. Try again!")
			continue
		}
		if guess != '\n' {
			if err := p.issueGuess(guess); err != nil {
				return errors.New("No hangman service detected. Bailing out!")
			}
		}
		if turnsLeft != p.game.Tally.TurnsLeft {
			turnsLeft, goodGuess = p.game.Tally.TurnsLeft, false
			continue
		}
		goodGuess = true
	}

	return nil
}

func (p *Player) newGame() (err error) {
	return p.call(context.TODO(), "GET", "new_game", nil)
}

func (p *Player) call(ctx context.Context, method, path string, payload io.Reader) error {
	url := urlFor(p.address, path)
	req, err := http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("boom: remote call %s crapped out!: %w", url, err)
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("boom: url call `%s failed with code (%d)", url, resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(p.game)
}

func (p *Player) gameOver() bool {
	switch p.game.Tally.Status {
	case Won:
		fmt.Print("\n👏  Noace!! You've just won\n\n")
		return true
	case Lost:
		fmt.Printf("\n😿  Meow! You've just lost. It was `%s\n\n", p.game.Word)
		return true
	case AlreadyGuessed, Active:
		return false
	default:
		return false
	}
}

func (p *Player) issueGuess(guess rune) error {
	body := struct {
		Game  *Game `json:"game"`
		Guess rune  `json:"guess"`
	}{
		Game:  p.game,
		Guess: guess,
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return p.call(context.TODO(), "POST", "guess", bytes.NewReader(payload))
}

func (p *Player) prompt(good bool) (rune, error) {
	ic := goodGuess
	if !good {
		ic = badGuess
	}
	fmt.Printf("\n%s  %10s [%d/%d]? ", ic, "Your Guess", p.game.Tally.TurnsLeft, MaxGuesses)
	char, _, err := bufio.NewReader(os.Stdin).ReadRune()
	if err != nil {
		return badEntry, err
	}

	return char, nil
}

// Helpers...

func urlFor(base, path string) string {
	return "http://" + base + "/api/v1/" + path
}
