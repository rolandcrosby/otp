# otp hash thingy

your two-factor auth code has gotta be all zeroes at some point, right? well, this little program tells you when that will happen. the first argument is the [totp:// URL](https://github.com/google/google-authenticator/wiki/Key-Uri-Format) of your two-factor token (or just its secret). the second argument can be either an exact six-digit code to search for, or the number of zeroes the code needs to begin with.

```
$ ./otp 'otpauth://totp/GitHub:rolandcrosby?secret=pbbbbbbbbbbbbbbt&issuer=GitHub' 3
code will be 000753 at 2020-02-19 02:44:42 UTC

$ ./otp pbbbbbbbbbbbbbbt 5
code will be 000008 at 2020-02-22 13:21:39 UTC

$ ./otp pbbbbbbbbbbbbbbt 555555
code will be 555555 at 2021-06-05 06:33:06 UTC
```