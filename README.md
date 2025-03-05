# xor-brute-force

This is a little toy program that brute forces stream ciphers with repeating keystreams.

![Tests](https://github.com/moritzhoeppner/xor-brute-force/actions/workflows/tests.yml/badge.svg)

## Usage

- Build the project with `make`
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
- Copy ciphertexts that were generated with the same keystream in one directory.
- Execute `bin/xor-brute-force -dist {path of your JSON file} -messages {directory with encrypted messages}`

You need at least two ciphertexts. However, the more you have, the better the results will be. The
program will only try to decrypt initial segments of the length of the shortest message.

## Example

Run `example/generate-example-messages.sh` to fill example/messages with 10 files, which contain
AES-256-CTR ciphertexts of English plaintexts. The encryption uses the same IV each time, which
means the counter blocks and hence the keystreams are the same for each file. `example/dist.json`
contains a frequency distribution for English texts I generated on the basis of texts from the
English Wikipedia. When you run
`bin/xor-brute-force -dist example/dist.json -messages ./example/messages`, you'll see the plaintexts
with only a few errors.

## How does it work?

If you re-use nonces/IVs for the encryption with stream ciphers, you'll get ciphertexts such that:
```
{ciphertext 1} XOR {ciphertext 2} = {plaintext 1} XOR {plaintext 2}
```
If you have enough of these ciphertexts, you can decrypt them byte by byte. First, calculate
```
{xored n} = {ciphertext 1} XOR {ciphertext n} (n > 1)
```
Now, guess the first byte of plaintext 1 (say, `b`). If the guess is correct, you'll get the
first byte of plaintext `n` by XORing `b` with the first byte of `{xored n}`. Calculate and save
these bytes for all `n`> 1. Now, guess a different first byte and do the same. When you tried all
possible values for `b`, choose the one that produces a result that approximates best the expected
frequency distribution of bytes.

This idea to essentially reduce the problem to a one-byte XOR cipher is taken from [this talk by
Thomas H. Ptacek](https://vimeo.com/41116595), which is a little dated but still very interesting.
