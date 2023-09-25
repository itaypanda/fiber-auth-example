# :sparkles: Go Fiber Auth Example

Authentication example for go-fiber

## :hash: Table of Contents

- [:sparkles: Go Fiber Auth Example](#sparkles-go-fiber-auth-example)
  - [:hash: Table of Contents](#hash-table-of-contents)
  - [:memo: Usage](#memo-usage)
  - [:world\_map: Routes](#world_map-routes)
    - [:lock:  Private Routes (Authenticated Only)](#lock--private-routes-authenticated-only)
    - [:unlock Public Rotues (Authenticated/Not Authenticated)](#unlock-public-rotues-authenticatednot-authenticated)
  - [:wave: Contact/Contribute](#wave-contactcontribute)
  - [:page\_with\_curl: License](#page_with_curl-license)
  - [:warning: Disclaimer](#warning-disclaimer)
  - [💙 Thanks](#-thanks)

## :memo: Usage

Copy [`.env.example`](./.env.example) to `.env` and fill in the required fields:

```bash
$ cp .env.example .env
```

```env
DATABASE_URL=postgres://username:password@localhost:5432/database
JWT_SECRET=REPLACE_ME_WITH_A_SECURE_LONG_SECRET_PLEASE
JWT_LIFESPAN=24
PORT=3000
BCRYPT_ROUNDS=12
```

**Please use a secure, long JWT secret!**
**If you don't set a JWT secret, one will be temporarily generated for you! IT DOES NOT SAVE!**

Start the server:

```bash
$ go run .
```

That's it, the server is up & working!

## :world_map: Routes

### :lock:  Private Routes (Authenticated Only)

**Aboutme Route ``[GET] /api/aboutme``:**

Takes the bearer token (from signin):
```
Bearer: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU3NTg3NDUsImlhdCI6MTY5NTY3MjM0NSwiaWQiOiI4OWIxNDU1Yy04YjYwLTQ5OGEtODMyNC03YjEyNzc3YzE5MmUifQ.V_QL64NusmThL7p16EUAEWx1BMIll22J0s9EF6_kY3c
```

And outputs (on success):
*[200 Ok]*
```json
{
  "email": "hi@example.com",
  "name": "Itaypanda"
}
```

### :unlock: Public Rotues (Authenticated/Not Authenticated)

**Signup Route ``[POST] /api/auth/signup``:**

Takes the body:
```json
{
  "name": "Itaypanda",        // Alphabetic (a-z A-Z) only, 3-16 chars
  "email": "hi@example.com",  // EMail
  "password": "password1234", // 8-256 chars
}
```

And outputs (on success):
*[201 Created]*
```json
{
  "id": "89b1455c-8b60-498a-8324-7b12777c192ey",
  "message": "User created"
}
```

**Signin Route ``[POST] /api/auth/signin``:**

Takes the body:
```json
{
  "email": "hi@example.com",  // EMail
  "password": "password1234", // 8-256 chars
}
```

And outputs (on success):
*[200 Ok]*
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU3NTg3NDUsImlhdCI6MTY5NTY3MjM0NSwiaWQiOiI4OWIxNDU1Yy04YjYwLTQ5OGEtODMyNC03YjEyNzc3YzE5MmUifQ.V_QL64NusmThL7p16EUAEWx1BMIll22J0s9EF6_kY3c"
}
```

## :wave: Contact/Contribute

Open an issue and I'll answer as soon as I'm available.

Feel free to contribute :D

Feel free to open an issue if you find any, or if you want to add a feature!


## :page_with_curl: License

This project is licensed under the [MIT license.](./LICENSE)

>Copyright (c) 2023 itaypanda
>
> Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

## :warning: Disclaimer

I am not a security expert, so this project may contain unsafe/not updated code. Please do your own research weather or not this project fits your needs, and weather it is up-to-date. This is just an example repository for learning purposes.

## 💙 Thanks

**Thanks to Tailwind CSS / Tailwind Labs, Inc.**

https://github.com/tailwindlabs/tailwindcss/blob/master/LICENSE

https://github.com/tailwindlabs/tailwindcs

https://tailwindcss.com/


**Thanks to Go Fiber / Fenny and Contributors**

https://github.com/gofiber/fiber/blob/master/LICENSE

https://github.com/gofiber/fiber

https://gofiber.io/


**Thanks to Go Dot Env / John Barton**

https://github.com/joho/godotenv/blob/main/LICENCE

https://github.com/joho/godotenv

http://godoc.org/github.com/joho/godotenv

**Thanks to GORM / Jinzhu**

https://github.com/go-gorm/gorm

https://github.com/go-gorm/gorm/blob/master/LICENSE

https://gorm.io/
