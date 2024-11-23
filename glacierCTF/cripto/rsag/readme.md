## Rivest–Shamir–Adleman-Germain

> My friend Sophie recently told me about this cool encryption algorithm. However she is not sure if it is secure. Can
> you help her by breaking it?

Given a ciphertext CT using an RSA system with four prime factors,
when modulus N=p*q*r*s, and the public key (e, N) and the ciphertext CT
are provided, decrypt the ciphertext using the `Chinese Remainder Theorem`
to simplify decryption.

We can find the prime factors of N using
https://www.alpertron.com.ar/ECM.HTM

Now the prime factors `p,q,r,s` are known
we can compute the partial private keys:

> d_i = e<sup>-1</sup> (mod &Phi;(pr) )

where i ={1,2,3,4} and pr = {p,q,r,s}

Decryption works using the formula:
> m = crt_result<sup>d1</sup> (mod p)

Since CRT has reduced the problem to the
p-modular space, it’s sufficient to use
`d1` (the private key for p) and `p` for decryption

flag `gctf{54dly_50ph13_63rm41n_pr1m35_wh3r3_n07_u53d_53curly}`