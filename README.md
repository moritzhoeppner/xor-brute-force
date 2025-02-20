# xor-brute-force

This is a little toy program that brute forces stream ciphers with repeating keystreams.

## Usage

- Build the project with `go build .`
- Describe the expected frequency distribution of plaintext bytes in a JSON file. For example, if
the byte 65 (in decimal representation) has a probability of 50% and the bytes 66 and 67 have
probabilities of 25%, respectively, your JSON file would be:
```json
{
  "65": 0.5,
  "66": 0.25,
  "67": 0.25
}
```
- Copy the encrypted messages in one directory.
- Execute `xor-brute-force -dist {path of your JSON file} -messages {directory with encrypted messages}`

## Example

Run `example/generate-example-messages.sh` to fill example/messages with 10 files, which contain
AES-256-CTR ciphertexts of English plaintexts. The encryption uses the same IV each time, which
means the counter blocks and hence the keystreams are the same for each file. `example/dist.json`
contains a frequency distribution for English texts I generated on the basis of texts from the
English Wikipedia. When you run
`xor-brute-force -dist example/dist.json -messages ./example/messages`, you'll see the plaintexts
with only a few errors.
