# otp hash thingy

your two-factor auth code has gotta be all zeroes at some point, right? well, this little program tells you when that will happen. the first argument is the [totp:// URL](https://github.com/google/google-authenticator/wiki/Key-Uri-Format) of your two-factor token (or just its secret). the second argument specifies what codes to look for:
- an exact six-digit code to search for
- the number of zeroes the code needs to begin with
- `duplicate`, which will let you know when the code will be the same for two time steps in a row

```
$ ./otp 'otpauth://totp/GitHub:rolandcrosby?secret=pbbbbbbbbbbbbbbt&issuer=GitHub' 3
code will be 000753 at 2020-02-19 02:44:42 UTC

$ ./otp pbbbbbbbbbbbbbbt 5
code will be 000008 at 2020-02-22 13:21:39 UTC

$ ./otp pbbbbbbbbbbbbbbt 555555
code will be 555555 at 2021-06-05 06:33:06 UTC

$ ./otp pbbbbbbbbbbbbbbt duplicate
code will be 734494 at 2022-09-27 22:41:25 UTC
code will be 734494 at 2022-09-27 22:41:55 UTC
```