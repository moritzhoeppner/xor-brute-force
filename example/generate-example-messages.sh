#!/bin/bash

# The plaintext samples are copied from
# https://en.wikipedia.org/w/index.php?title=Crime_and_Punishment_(1983_film)&oldid=1266114572
plain_texts=(
  "A slaughterhouse worker, Antti Rahikainen, murders a man."
  "finds the woman who witnessed the murder, Eeva Laakso, an"
  "Lining Hostel. He arrives home and finds that he's been s"
  "ime together. She finds the gun at his apartment and pock"
  "businessman who had been murdered. It is revealed that th"
  " was killed by Honkanen in a drink-driving crash, and whe"
  "him, Rahikainen swore to take revenge. Eeva arrives at th"
  " not identify Rahikainen as the murderer, forcing the cop"
  "on a ferry at night, where he tells her he killed Honkane"
  "he \"found him disgusting\" and \"wanted to show that things"
)

mkdir -p example/messages
counter=0
for text in "${plain_texts[@]}"; do
  echo -n "$text" | openssl enc -aes-256-ctr \
    -K 000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f \
    -iv 00000000000000000000000000000000 \
    -out example/messages/$counter.enc
  counter=$((counter + 1))
done

