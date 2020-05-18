<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Reflection Lab...

## <img src="../assets/lab.png" width="auto" height="32"/> Your Mission

> Use reflection to compute a BookInfo summary instance.

1. Clone the [Labs Repo](https://github.com/gopherland/labs2)
2. Cd reflection
3. Implement an hydrate function to populate a BookInfo details.
   1. A BookInfo type contains the following fields:
      1. Book represents a book file on disk
      2. Lines represents the number of lines in the book
      3. Words represents the number of words in the book
      4. IBN represents a unique book ID computed either via sha1 or md5
   2. The IBN field contains a struct tag to indicate the method to use when computing the IBN. These can be either `sha1` or `md5`
4. The hydrate function will take in a partially filled BookInfo containing the Book field pre-populated with the book location on disk. Using reflection, compute the rest of the BookInfo fields.
   1. Hydrate must fill in the following fields:
      1. Read the book file and compute the number of lines and words.
      2. Set the corresponding fields on the BookInfo instance
      3. Next read the struct tag on the BookInfo.IBN field to determine the IBN compute method. Based on your findings compute and set the IBN field using either sha1 or md5
5. Call your hydrate function and print the resulting BookInfo.

## Sample Outputs

### Given

```go
type BookInfo struct {
   Book  string
   Lines int
   IBN   string `md5:"ibn"`
   Words int
}
b := BookInfo{
   File: "assets/100west.txt,
}
```

### Output

```text
{
   assets/100west.txt
   416
   a8d80d7229a40ce1a69c4a042de269fc # => md5 of book file name
   5650
}
```

### Given

```go
type BookInfo struct {
   Book  string
   Lines int
   IBN   string `sha1:"ibn"`
   Words int
}
b := BookInfo{
   File: "assets/100west.txt,
}
```

### Output

```text
{
   assets/100west.txt
   416
   418344303f2d5611ab2e70abcb038dedeb5e42f0 # => sha1 of book file name
   5650
}
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
