# My Project

Base URLs:

# Authentication

-   HTTP Authentication, scheme: bearer

# Default

## POST Register

POST /register

> Body Parameters

```json
{
    "FullName": "Ananda Ravi Kuntadi",
    "NoPhone": "434434",
    "EmailAddress": "beratstress39@gmail.com",
    "Password": "alo123",
    "ConfirmPassword": "alo123"
}
```

### Params

| Name              | Location | Type   | Required | Description |
| ----------------- | -------- | ------ | -------- | ----------- |
| body              | body     | object | no       | none        |
| » FullName        | body     | string | yes      | none        |
| » NoPhone         | body     | string | yes      | none        |
| » EmailAddress    | body     | string | yes      | none        |
| » Password        | body     | string | yes      | none        |
| » ConfirmPassword | body     | string | yes      | none        |

> Response Examples

> Register

```json
{
    "status": "Verification email has been sent. Please verify your account."
}
```

### Responses

| HTTP Status Code | Meaning                                                      | Description | Data schema |
| ---------------- | ------------------------------------------------------------ | ----------- | ----------- |
| 201              | [Created](https://tools.ietf.org/html/rfc7231#section-6.3.2) | Register    | Inline      |

### Responses Data Schema

HTTP Status Code **201**

| Name     | Type   | Required | Restrictions | Title | description |
| -------- | ------ | -------- | ------------ | ----- | ----------- |
| » status | string | true     | none         |       | none        |

## POST Resend Verification

POST /resend-verification

> Body Parameters

```json
{
    "email_address": "beratstress39@gmail.com"
}
```

### Params

| Name            | Location | Type   | Required | Description |
| --------------- | -------- | ------ | -------- | ----------- |
| body            | body     | object | no       | none        |
| » email_address | body     | string | yes      | none        |

> Response Examples

> Resend

```json
{
    "status": "Verification email has been resent. Please verify your account within 10 minutes."
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Resend      | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name     | Type   | Required | Restrictions | Title | description |
| -------- | ------ | -------- | ------------ | ----- | ----------- |
| » status | string | true     | none         |       | none        |

## POST Verify Code

POST /verify

> Body Parameters

```json
{
    "code": "6036"
}
```

### Params

| Name   | Location | Type   | Required | Description |
| ------ | -------- | ------ | -------- | ----------- |
| body   | body     | object | no       | none        |
| » code | body     | string | yes      | none        |

> Response Examples

> Verify Code

```json
{
    "status": "User verified successfully!",
    "user": {
        "id": 10,
        "full_name": "Ananda Ravi Kuntadi",
        "no_phone": "434434",
        "email_address": "beratstress39@gmail.com",
        "is_verified": true
    }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Verify Code | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name             | Type    | Required | Restrictions | Title | description |
| ---------------- | ------- | -------- | ------------ | ----- | ----------- |
| » status         | string  | true     | none         |       | none        |
| » user           | object  | true     | none         |       | none        |
| »» id            | integer | true     | none         |       | none        |
| »» full_name     | string  | true     | none         |       | none        |
| »» no_phone      | string  | true     | none         |       | none        |
| »» email_address | string  | true     | none         |       | none        |
| »» is_verified   | boolean | true     | none         |       | none        |

## POST Login

POST /login

> Body Parameters

```json
{
    "EmailAddress": "arknkoo6@gmail.com",
    "Password": "alo123"
}
```

### Params

| Name           | Location | Type   | Required | Description |
| -------------- | -------- | ------ | -------- | ----------- |
| body           | body     | object | no       | none        |
| » EmailAddress | body     | string | yes      | none        |
| » Password     | body     | string | yes      | none        |

> Response Examples

> Login

```json
{
    "status": "User logged in successfully!",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ3MTA5MTEsInVzZXJJRCI6Mn0.djBYiI7S1xr9BeExd1gohpwKT1hOv4OHZ6NyM0OXlK0",
    "user": {
        "ID": 2,
        "CreatedAt": "2024-01-05T17:30:59.482+07:00",
        "UpdatedAt": "2024-01-05T17:30:59.482+07:00",
        "DeletedAt": null,
        "FullName": "Ananda Ravi Kuntadi",
        "NoPhone": "0812",
        "EmailAddress": "anandaravik@gmail.com",
        "Password": "$2a$10$ASCXnSEFAHfyLeAftA21ru3vTu.9SKpvJL4UESj0v9EOLSGJiwl8q",
        "ConfirmPassword": ""
    }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Login       | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name               | Type    | Required | Restrictions | Title | description |
| ------------------ | ------- | -------- | ------------ | ----- | ----------- |
| » status           | string  | true     | none         |       | none        |
| » token            | string  | true     | none         |       | none        |
| » user             | object  | true     | none         |       | none        |
| »» ID              | integer | true     | none         |       | none        |
| »» CreatedAt       | string  | true     | none         |       | none        |
| »» UpdatedAt       | string  | true     | none         |       | none        |
| »» DeletedAt       | null    | true     | none         |       | none        |
| »» FullName        | string  | true     | none         |       | none        |
| »» NoPhone         | string  | true     | none         |       | none        |
| »» EmailAddress    | string  | true     | none         |       | none        |
| »» Password        | string  | true     | none         |       | none        |
| »» ConfirmPassword | string  | true     | none         |       | none        |

## POST Forgot Password

POST /forgot-password

> Body Parameters

```json
{
    "email_address": "arknkoo6@gmail.com"
}
```

### Params

| Name            | Location | Type   | Required | Description |
| --------------- | -------- | ------ | -------- | ----------- |
| body            | body     | object | no       | none        |
| » email_address | body     | string | yes      | none        |

> Response Examples

> Forgot Password

```json
{
    "status": "Verification email has been sent."
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description     | Data schema |
| ---------------- | ------------------------------------------------------- | --------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Forgot Password | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name     | Type   | Required | Restrictions | Title | description |
| -------- | ------ | -------- | ------------ | ----- | ----------- |
| » status | string | true     | none         |       | none        |

## POST Verify Forgot Password

POST /verify-token

> Body Parameters

```json
{
    "email_address": "arknkoo6@gmail.com",
    "verification_code": "3698"
}
```

### Params

| Name                | Location | Type   | Required | Description |
| ------------------- | -------- | ------ | -------- | ----------- |
| body                | body     | object | no       | none        |
| » email_address     | body     | string | yes      | none        |
| » verification_code | body     | string | yes      | none        |

> Response Examples

> Verify

```json
{
    "status": "Token verified successfully!",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Imtvb3JhdmkxMkBnbWFpbC5jb20iLCJleHAiOjE3MDQ4Mjk3OTJ9.2mWHBJcmViNiEegSe-ZwmXWC9NoVB-YY77lVbZObVKA"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Verify      | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name     | Type   | Required | Restrictions | Title | description |
| -------- | ------ | -------- | ------------ | ----- | ----------- |
| » status | string | true     | none         |       | none        |
| » token  | string | true     | none         |       | none        |

## POST Reset Password

POST /reset-password

> Body Parameters

```json
{
    "password": "STRESSMAN",
    "confirm_password": "STRESSMAN"
}
```

### Params

| Name               | Location | Type   | Required | Description |
| ------------------ | -------- | ------ | -------- | ----------- |
| body               | body     | object | no       | none        |
| » password         | body     | string | yes      | none        |
| » confirm_password | body     | string | yes      | none        |

> Response Examples

> Reset Password

```json
{
    "status": "Password reset successfully!"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description    | Data schema |
| ---------------- | ------------------------------------------------------- | -------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Reset Password | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name     | Type   | Required | Restrictions | Title | description |
| -------- | ------ | -------- | ------------ | ----- | ----------- |
| » status | string | true     | none         |       | none        |

## POST Create Profile

POST /profile

> Body Parameters

```json
{
    "NamaLengkap": "RAVI KUNTADI",
    "TempatLahir": "Yogyakarta",
    "tanggal_lahir": "2006-01-02",
    "Alamat": "Jl. Malioboro No.1",
    "Alergi": "Peanuts"
}
```

### Params

| Name            | Location | Type   | Required | Description |
| --------------- | -------- | ------ | -------- | ----------- |
| body            | body     | object | no       | none        |
| » NamaLengkap   | body     | string | yes      | none        |
| » TempatLahir   | body     | string | yes      | none        |
| » tanggal_lahir | body     | string | yes      | none        |
| » Alamat        | body     | string | yes      | none        |
| » Alergi        | body     | string | yes      | none        |

> Response Examples

> Create Profile

```json
{
    "message": "Profile created successfully"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description    | Data schema |
| ---------------- | ------------------------------------------------------- | -------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Create Profile | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

## GET Show Profile ALL

GET /profile

> Response Examples

> Show Profile ALL

```json
[
    {
        "ID": 1,
        "CreatedAt": "2024-01-09T22:08:04.355+07:00",
        "UpdatedAt": "2024-01-09T22:08:04.355+07:00",
        "DeletedAt": null,
        "UserID": 23,
        "NamaLengkap": "",
        "TempatLahir": "",
        "TanggalLahir": "2006-01-02T00:00:00+07:00",
        "Alamat": "Jl. Malioboro No.1",
        "Alergi": "Peanuts",
        "User": {
            "ID": 23,
            "CreatedAt": "2024-01-07T06:28:10.38+07:00",
            "UpdatedAt": "2024-01-07T06:28:10.38+07:00",
            "DeletedAt": null,
            "FullName": "Ananda Ravi Kuntadi",
            "NoPhone": "434434",
            "EmailAddress": "beratstress39@gmail.com",
            "Password": "$2a$10$zE5o5xQWolBii0gHCX3FSeLG13AaZnpDLuAyeTqBwY8ggCn2jwZeq",
            "ConfirmPassword": "",
            "VerificationCode": "6267",
            "IsVerified": true,
            "Profiles": null
        }
    },
    {
        "ID": 2,
        "CreatedAt": "2024-01-09T22:08:53.899+07:00",
        "UpdatedAt": "2024-01-09T22:08:53.899+07:00",
        "DeletedAt": null,
        "UserID": 23,
        "NamaLengkap": "RAVI KUNTADI",
        "TempatLahir": "Yogyakarta",
        "TanggalLahir": "2006-01-02T00:00:00+07:00",
        "Alamat": "Jl. Malioboro No.1",
        "Alergi": "Peanuts",
        "User": {
            "ID": 23,
            "CreatedAt": "2024-01-07T06:28:10.38+07:00",
            "UpdatedAt": "2024-01-07T06:28:10.38+07:00",
            "DeletedAt": null,
            "FullName": "Ananda Ravi Kuntadi",
            "NoPhone": "434434",
            "EmailAddress": "beratstress39@gmail.com",
            "Password": "$2a$10$zE5o5xQWolBii0gHCX3FSeLG13AaZnpDLuAyeTqBwY8ggCn2jwZeq",
            "ConfirmPassword": "",
            "VerificationCode": "6267",
            "IsVerified": true,
            "Profiles": null
        }
    },
    {
        "ID": 3,
        "CreatedAt": "2024-01-09T22:09:32.059+07:00",
        "UpdatedAt": "2024-01-09T22:09:32.059+07:00",
        "DeletedAt": null,
        "UserID": 23,
        "NamaLengkap": "RAVI KUNTADI",
        "TempatLahir": "Yogyakarta",
        "TanggalLahir": "2006-01-02T00:00:00+07:00",
        "Alamat": "Jl. Malioboro No.1",
        "Alergi": "Peanuts",
        "User": {
            "ID": 23,
            "CreatedAt": "2024-01-07T06:28:10.38+07:00",
            "UpdatedAt": "2024-01-07T06:28:10.38+07:00",
            "DeletedAt": null,
            "FullName": "Ananda Ravi Kuntadi",
            "NoPhone": "434434",
            "EmailAddress": "beratstress39@gmail.com",
            "Password": "$2a$10$zE5o5xQWolBii0gHCX3FSeLG13AaZnpDLuAyeTqBwY8ggCn2jwZeq",
            "ConfirmPassword": "",
            "VerificationCode": "6267",
            "IsVerified": true,
            "Profiles": null
        }
    }
]
```

### Responses

| HTTP Status Code | Meaning                                                 | Description      | Data schema |
| ---------------- | ------------------------------------------------------- | ---------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Show Profile ALL | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name                | Type    | Required | Restrictions | Title | description |
| ------------------- | ------- | -------- | ------------ | ----- | ----------- |
| » ID                | integer | true     | none         |       | none        |
| » CreatedAt         | string  | true     | none         |       | none        |
| » UpdatedAt         | string  | true     | none         |       | none        |
| » DeletedAt         | null    | true     | none         |       | none        |
| » UserID            | integer | true     | none         |       | none        |
| » NamaLengkap       | string  | true     | none         |       | none        |
| » TempatLahir       | string  | true     | none         |       | none        |
| » TanggalLahir      | string  | true     | none         |       | none        |
| » Alamat            | string  | true     | none         |       | none        |
| » Alergi            | string  | true     | none         |       | none        |
| » User              | object  | true     | none         |       | none        |
| »» ID               | integer | true     | none         |       | none        |
| »» CreatedAt        | string  | true     | none         |       | none        |
| »» UpdatedAt        | string  | true     | none         |       | none        |
| »» DeletedAt        | null    | true     | none         |       | none        |
| »» FullName         | string  | true     | none         |       | none        |
| »» NoPhone          | string  | true     | none         |       | none        |
| »» EmailAddress     | string  | true     | none         |       | none        |
| »» Password         | string  | true     | none         |       | none        |
| »» ConfirmPassword  | string  | true     | none         |       | none        |
| »» VerificationCode | string  | true     | none         |       | none        |
| »» IsVerified       | boolean | true     | none         |       | none        |
| »» Profiles         | null    | true     | none         |       | none        |

## GET Show Profile

GET /profile/1

> Response Examples

> Show Profil

```json
{
    "ID": 3,
    "CreatedAt": "2024-01-09T22:09:32.059+07:00",
    "UpdatedAt": "2024-01-09T22:09:32.059+07:00",
    "DeletedAt": null,
    "UserID": 23,
    "NamaLengkap": "RAVI KUNTADI",
    "TempatLahir": "Yogyakarta",
    "TanggalLahir": "2006-01-02T00:00:00+07:00",
    "Alamat": "Jl. Malioboro No.1",
    "Alergi": "Peanuts",
    "User": {
        "ID": 23,
        "CreatedAt": "2024-01-07T06:28:10.38+07:00",
        "UpdatedAt": "2024-01-07T06:28:10.38+07:00",
        "DeletedAt": null,
        "FullName": "Ananda Ravi Kuntadi",
        "NoPhone": "434434",
        "EmailAddress": "beratstress39@gmail.com",
        "Password": "$2a$10$zE5o5xQWolBii0gHCX3FSeLG13AaZnpDLuAyeTqBwY8ggCn2jwZeq",
        "ConfirmPassword": "",
        "VerificationCode": "6267",
        "IsVerified": true,
        "Profiles": null
    }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Show Profil | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name                | Type    | Required | Restrictions | Title | description |
| ------------------- | ------- | -------- | ------------ | ----- | ----------- |
| » ID                | integer | true     | none         |       | none        |
| » CreatedAt         | string  | true     | none         |       | none        |
| » UpdatedAt         | string  | true     | none         |       | none        |
| » DeletedAt         | null    | true     | none         |       | none        |
| » UserID            | integer | true     | none         |       | none        |
| » NamaLengkap       | string  | true     | none         |       | none        |
| » TempatLahir       | string  | true     | none         |       | none        |
| » TanggalLahir      | string  | true     | none         |       | none        |
| » Alamat            | string  | true     | none         |       | none        |
| » Alergi            | string  | true     | none         |       | none        |
| » User              | object  | true     | none         |       | none        |
| »» ID               | integer | true     | none         |       | none        |
| »» CreatedAt        | string  | true     | none         |       | none        |
| »» UpdatedAt        | string  | true     | none         |       | none        |
| »» DeletedAt        | null    | true     | none         |       | none        |
| »» FullName         | string  | true     | none         |       | none        |
| »» NoPhone          | string  | true     | none         |       | none        |
| »» EmailAddress     | string  | true     | none         |       | none        |
| »» Password         | string  | true     | none         |       | none        |
| »» ConfirmPassword  | string  | true     | none         |       | none        |
| »» VerificationCode | string  | true     | none         |       | none        |
| »» IsVerified       | boolean | true     | none         |       | none        |
| »» Profiles         | null    | true     | none         |       | none        |

## PUT Update Profile

PUT /profile/1

> Body Parameters

```json
{
    "NamaLengkap": "Rapi Kuntadi"
}
```

### Params

| Name          | Location | Type   | Required | Description |
| ------------- | -------- | ------ | -------- | ----------- |
| body          | body     | object | no       | none        |
| » NamaLengkap | body     | string | yes      | none        |

> Response Examples

> Update Profile

```json
{
    "message": "Profile updated successfully"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description    | Data schema |
| ---------------- | ------------------------------------------------------- | -------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Update Profile | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

## GET Show Profile

GET /profile/3

> Response Examples

> Show Profil

```json
{
    "ID": 3,
    "CreatedAt": "2024-01-09T22:09:32.059+07:00",
    "UpdatedAt": "2024-01-09T22:09:32.059+07:00",
    "DeletedAt": null,
    "UserID": 23,
    "NamaLengkap": "RAVI KUNTADI",
    "TempatLahir": "Yogyakarta",
    "TanggalLahir": "2006-01-02T00:00:00+07:00",
    "Alamat": "Jl. Malioboro No.1",
    "Alergi": "Peanuts",
    "User": {
        "ID": 23,
        "CreatedAt": "2024-01-07T06:28:10.38+07:00",
        "UpdatedAt": "2024-01-07T06:28:10.38+07:00",
        "DeletedAt": null,
        "FullName": "Ananda Ravi Kuntadi",
        "NoPhone": "434434",
        "EmailAddress": "beratstress39@gmail.com",
        "Password": "$2a$10$zE5o5xQWolBii0gHCX3FSeLG13AaZnpDLuAyeTqBwY8ggCn2jwZeq",
        "ConfirmPassword": "",
        "VerificationCode": "6267",
        "IsVerified": true,
        "Profiles": null
    }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Show Profil | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name                | Type    | Required | Restrictions | Title | description |
| ------------------- | ------- | -------- | ------------ | ----- | ----------- |
| » ID                | integer | true     | none         |       | none        |
| » CreatedAt         | string  | true     | none         |       | none        |
| » UpdatedAt         | string  | true     | none         |       | none        |
| » DeletedAt         | null    | true     | none         |       | none        |
| » UserID            | integer | true     | none         |       | none        |
| » NamaLengkap       | string  | true     | none         |       | none        |
| » TempatLahir       | string  | true     | none         |       | none        |
| » TanggalLahir      | string  | true     | none         |       | none        |
| » Alamat            | string  | true     | none         |       | none        |
| » Alergi            | string  | true     | none         |       | none        |
| » User              | object  | true     | none         |       | none        |
| »» ID               | integer | true     | none         |       | none        |
| »» CreatedAt        | string  | true     | none         |       | none        |
| »» UpdatedAt        | string  | true     | none         |       | none        |
| »» DeletedAt        | null    | true     | none         |       | none        |
| »» FullName         | string  | true     | none         |       | none        |
| »» NoPhone          | string  | true     | none         |       | none        |
| »» EmailAddress     | string  | true     | none         |       | none        |
| »» Password         | string  | true     | none         |       | none        |
| »» ConfirmPassword  | string  | true     | none         |       | none        |
| »» VerificationCode | string  | true     | none         |       | none        |
| »» IsVerified       | boolean | true     | none         |       | none        |
| »» Profiles         | null    | true     | none         |       | none        |

## PUT Update Profile

PUT /profile/3

> Body Parameters

```json
{
    "NamaLengkap": "Rapi Kuntadi"
}
```

### Params

| Name          | Location | Type   | Required | Description |
| ------------- | -------- | ------ | -------- | ----------- |
| body          | body     | object | no       | none        |
| » NamaLengkap | body     | string | yes      | none        |

> Response Examples

> Update Profile

```json
{
    "message": "Profile updated successfully"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description    | Data schema |
| ---------------- | ------------------------------------------------------- | -------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Update Profile | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

## DELETE Delete Profile

DELETE /profile/2

> Response Examples

> Delete Profile

```json
{
    "message": "Profile deleted successfully"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description    | Data schema |
| ---------------- | ------------------------------------------------------- | -------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Delete Profile | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

## POST Complete Task Timenya NOW

POST /profile/3/task/4/complete

> Response Examples

> Complete Task

```json
{
    "message": "Task completed successfully"
}
```

```json
{
    "error": "this task has already been completed today"
}
```

### Responses

| HTTP Status Code | Meaning                                                                    | Description   | Data schema |
| ---------------- | -------------------------------------------------------------------------- | ------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)                    | Complete Task | Inline      |
| 500              | [Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1) | Complete Task | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

HTTP Status Code **500**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » error | string | true     | none         |       | none        |

## GET Show Task Each Day

GET /profile/3/tasks

### Params

| Name | Location | Type   | Required | Description |
| ---- | -------- | ------ | -------- | ----------- |
| date | query    | string | yes      | none        |

> Response Examples

> Show Task Each Day

```json
[
    {
        "Bulan": 1,
        "Nama": "Brush teeth after breakfast",
        "Profile ID": 3,
        "Status": true,
        "Tahun": 2024,
        "Tanggal": 10,
        "Task ID": 1
    },
    {
        "Bulan": 1,
        "Nama": "Brush teeth for 2 minutes",
        "Profile ID": 3,
        "Status": true,
        "Tahun": 2024,
        "Tanggal": 10,
        "Task ID": 2
    },
    {
        "Bulan": 1,
        "Nama": "Brush teeth before sleep",
        "Profile ID": 3,
        "Status": true,
        "Tahun": 2024,
        "Tanggal": 10,
        "Task ID": 3
    },
    {
        "Bulan": 1,
        "Nama": "Floss once a day",
        "Profile ID": 3,
        "Status": true,
        "Tahun": 2024,
        "Tanggal": 10,
        "Task ID": 4
    },
    {
        "Bulan": 1,
        "Nama": "Use fluoride toothpaste",
        "Profile ID": 3,
        "Status": true,
        "Tahun": 2024,
        "Tanggal": 10,
        "Task ID": 5
    }
]
```

### Responses

| HTTP Status Code | Meaning                                                 | Description        | Data schema |
| ---------------- | ------------------------------------------------------- | ------------------ | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Show Task Each Day | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name         | Type    | Required | Restrictions | Title | description |
| ------------ | ------- | -------- | ------------ | ----- | ----------- |
| » Bulan      | integer | true     | none         |       | none        |
| » Nama       | string  | true     | none         |       | none        |
| » Profile ID | integer | true     | none         |       | none        |
| » Status     | boolean | true     | none         |       | none        |
| » Tahun      | integer | true     | none         |       | none        |
| » Tanggal    | integer | true     | none         |       | none        |
| » Task ID    | integer | true     | none         |       | none        |

## PUT Undo Task Timenya NOW

PUT /profile/3/task/5/undo

> Response Examples

> Undo Task

```json
{
    "message": "Task has been marked as not completed"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Undo Task   | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

## GET All Task

GET /tasks

> Response Examples

> All Task

```json
[
    {
        "ID": 1,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "Name": "Brush teeth after breakfast",
        "Points": 5
    },
    {
        "ID": 2,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "Name": "Brush teeth for 2 minutes",
        "Points": 5
    },
    {
        "ID": 3,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "Name": "Brush teeth before sleep",
        "Points": 5
    },
    {
        "ID": 4,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "Name": "Floss once a day",
        "Points": 5
    },
    {
        "ID": 5,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "Name": "Use fluoride toothpaste",
        "Points": 5
    }
]
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | All Task    | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name        | Type    | Required | Restrictions | Title | description |
| ----------- | ------- | -------- | ------------ | ----- | ----------- |
| » ID        | integer | true     | none         |       | none        |
| » CreatedAt | string  | true     | none         |       | none        |
| » UpdatedAt | string  | true     | none         |       | none        |
| » DeletedAt | null    | true     | none         |       | none        |
| » Name      | string  | true     | none         |       | none        |
| » Points    | integer | true     | none         |       | none        |

## GET Recap

GET /profile/3/completedTasks

> Response Examples

> Recap

```json
{
    "data": [
        {
            "DayCompleted": 0,
            "TotalPoints": 0,
            "completedTasks": 0,
            "periode": "Jul - Dec 2024"
        },
        {
            "DayCompleted": 3,
            "TotalPoints": 5,
            "completedTasks": 11,
            "periode": "Jan - Jun 2024"
        }
    ]
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Recap       | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name              | Type     | Required | Restrictions | Title | description |
| ----------------- | -------- | -------- | ------------ | ----- | ----------- |
| » data            | [object] | true     | none         |       | none        |
| »» DayCompleted   | integer  | true     | none         |       | none        |
| »» TotalPoints    | integer  | true     | none         |       | none        |
| »» completedTasks | integer  | true     | none         |       | none        |
| »» periode        | string   | true     | none         |       | none        |

## POST Add Question

POST /profile/3/question

> Body Parameters

```json
{
    "Tag": "Caries",
    "Question": "What are the symptoms of tooth decay?"
}
```

### Params

| Name       | Location | Type   | Required | Description |
| ---------- | -------- | ------ | -------- | ----------- |
| body       | body     | object | no       | none        |
| » Tag      | body     | string | yes      | none        |
| » Question | body     | string | yes      | none        |

> Response Examples

> New Request

```json
{
    "message": "Question created successfully"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | New Request | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

## POST Add Question

POST /profile/1/question

> Body Parameters

```json
{
    "Tag": "Caries",
    "Question": "What are the symptoms of tooth decay?"
}
```

### Params

| Name       | Location | Type   | Required | Description |
| ---------- | -------- | ------ | -------- | ----------- |
| body       | body     | object | no       | none        |
| » Tag      | body     | string | yes      | none        |
| » Question | body     | string | yes      | none        |

> Response Examples

> New Request

```json
{
    "message": "Question created successfully"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | New Request | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

## GET Get Question ALL

GET /questions

> Response Examples

> Get Question ALL

```json
[
    {
        "ID": 2,
        "CreatedAt": "2024-01-12T17:14:51.347+07:00",
        "UpdatedAt": "2024-01-12T18:30:00.986+07:00",
        "DeletedAt": null,
        "Tag": "Caries",
        "Question": "What are the symptoms of tooth decay?",
        "Answer": "The symptoms of tooth decay can vary depending on the severity of the decay. In the early stages, tooth decay may not cause any symptoms at all. However, as the decay progresses, you may begin to experience the following symptoms: Sensitivity to hot or cold foods and drinks Pain when biting or chewingLoose or chipped teeth Dark spots or cavities on the teeth If you experience any of these symptoms, it is important to see a dentist as soon as possible. Tooth decay can be a serious condition that can lead to tooth loss if left untreated.",
        "AnsweredAt": "2024-01-12T18:30:00.792+07:00",
        "ProfileID": 3,
        "DentistID": 2,
        "NamaLengkap": "Rapi Kuntadi"
    },
    {
        "ID": 3,
        "CreatedAt": "2024-01-12T17:16:05.863+07:00",
        "UpdatedAt": "2024-01-12T18:26:17.785+07:00",
        "DeletedAt": null,
        "Tag": "Cleaning",
        "Question": "How often should I brush my teeth?",
        "Answer": "",
        "AnsweredAt": null,
        "ProfileID": 3,
        "DentistID": null,
        "NamaLengkap": "Rapi Kuntadi"
    },
    {
        "ID": 4,
        "CreatedAt": "2024-01-12T17:16:18.18+07:00",
        "UpdatedAt": "2024-01-12T17:16:18.18+07:00",
        "DeletedAt": null,
        "Tag": "Fluoride",
        "Question": "How does fluoride help teeth?",
        "Answer": "",
        "AnsweredAt": null,
        "ProfileID": 3,
        "DentistID": null,
        "NamaLengkap": "Rapi Kuntadi"
    },
    {
        "ID": 5,
        "CreatedAt": "2024-01-12T17:16:31.011+07:00",
        "UpdatedAt": "2024-01-12T17:16:31.011+07:00",
        "DeletedAt": null,
        "Tag": "Prevention",
        "Question": "How can I keep my teeth healthy?",
        "Answer": "",
        "AnsweredAt": null,
        "ProfileID": 3,
        "DentistID": null,
        "NamaLengkap": "Rapi Kuntadi"
    }
]
```

### Responses

| HTTP Status Code | Meaning                                                 | Description      | Data schema |
| ---------------- | ------------------------------------------------------- | ---------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Get Question ALL | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name          | Type         | Required | Restrictions | Title | description |
| ------------- | ------------ | -------- | ------------ | ----- | ----------- |
| » ID          | integer      | true     | none         |       | none        |
| » CreatedAt   | string       | true     | none         |       | none        |
| » UpdatedAt   | string       | true     | none         |       | none        |
| » DeletedAt   | null         | true     | none         |       | none        |
| » Tag         | string       | true     | none         |       | none        |
| » Question    | string       | true     | none         |       | none        |
| » Answer      | string       | true     | none         |       | none        |
| » AnsweredAt  | string¦null  | true     | none         |       | none        |
| » ProfileID   | integer      | true     | none         |       | none        |
| » DentistID   | integer¦null | true     | none         |       | none        |
| » NamaLengkap | string       | true     | none         |       | none        |

## GET Get Question By ID

GET /question/3

> Response Examples

> Get Question By ID

```json
{
    "ID": 3,
    "CreatedAt": "2024-01-12T17:16:05.863+07:00",
    "UpdatedAt": "2024-01-12T18:26:17.785+07:00",
    "DeletedAt": null,
    "Tag": "Cleaning",
    "Question": "How often should I brush my teeth?",
    "Answer": "",
    "AnsweredAt": null,
    "ProfileID": 3,
    "DentistID": null,
    "NamaLengkap": "Rapi Kuntadi"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description        | Data schema |
| ---------------- | ------------------------------------------------------- | ------------------ | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Get Question By ID | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name          | Type    | Required | Restrictions | Title | description |
| ------------- | ------- | -------- | ------------ | ----- | ----------- |
| » ID          | integer | true     | none         |       | none        |
| » CreatedAt   | string  | true     | none         |       | none        |
| » UpdatedAt   | string  | true     | none         |       | none        |
| » DeletedAt   | null    | true     | none         |       | none        |
| » Tag         | string  | true     | none         |       | none        |
| » Question    | string  | true     | none         |       | none        |
| » Answer      | string  | true     | none         |       | none        |
| » AnsweredAt  | null    | true     | none         |       | none        |
| » ProfileID   | integer | true     | none         |       | none        |
| » DentistID   | null    | true     | none         |       | none        |
| » NamaLengkap | string  | true     | none         |       | none        |

## DELETE Delete Question

DELETE /question/6

> Response Examples

> Delete Question

```json
{
    "message": "Question deleted successfully"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description     | Data schema |
| ---------------- | ------------------------------------------------------- | --------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Delete Question | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

## POST Add Dentist

POST /dentists

> Body Parameters

```json
{
    "Name": "Dr Molly Weasley",
    "Specialist": "Paediatric",
    "AboutMe": "I am a dentist at Saiful Anwar Hospital in Malang, Indonesia. I have been working here for three years, and I have treated hundreds of patients. I am passionate about dentistry, and I believe that everyone deserves to have healthy, beautiful teeth.",
    "WorkYearExperience": 5,
    "PatientCount": 200
}
```

### Params

| Name                 | Location | Type    | Required | Description |
| -------------------- | -------- | ------- | -------- | ----------- |
| body                 | body     | object  | no       | none        |
| » Name               | body     | string  | yes      | none        |
| » Specialist         | body     | string  | yes      | none        |
| » AboutMe            | body     | string  | yes      | none        |
| » WorkYearExperience | body     | integer | yes      | none        |
| » PatientCount       | body     | integer | yes      | none        |

> Response Examples

> New Request

```json
{
    "message": "Profile created successfully"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | New Request | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

## GET Get Dentist ALL

GET /dentists

> Body Parameters

```json
{}
```

### Params

| Name | Location | Type   | Required | Description |
| ---- | -------- | ------ | -------- | ----------- |
| body | body     | object | no       | none        |

> Response Examples

> Get Dentist ALL

```json
[
    {
        "ID": 1,
        "CreatedAt": "2024-03-01T01:10:51.263Z",
        "UpdatedAt": "2024-03-01T01:10:51.263Z",
        "DeletedAt": null,
        "Name": "Dr Molly Weasley",
        "Specialist": "Paediatric",
        "AboutMe": "I am a dentist at Saiful Anwar Hospital in Malang, Indonesia. I have been working here for three years, and I have treated hundreds of patients. I am passionate about dentistry, and I believe that everyone deserves to have healthy, beautiful teeth.",
        "WorkYearExperience": 5,
        "PatientCount": 200,
        "Picture": "",
        "OnlineConsultations": null,
        "OfflineConsultations": null,
        "Appointments": null,
        "Ratings": null,
        "Questions": null
    },
    {
        "ID": 2,
        "CreatedAt": "2024-03-01T03:02:50.423Z",
        "UpdatedAt": "2024-03-01T03:02:50.423Z",
        "DeletedAt": null,
        "Name": "Dr Molly Weasley",
        "Specialist": "Paediatric",
        "AboutMe": "I am a dentist at Saiful Anwar Hospital in Malang, Indonesia. I have been working here for three years, and I have treated hundreds of patients. I am passionate about dentistry, and I believe that everyone deserves to have healthy, beautiful teeth.",
        "WorkYearExperience": 5,
        "PatientCount": 200,
        "Picture": "",
        "OnlineConsultations": null,
        "OfflineConsultations": null,
        "Appointments": null,
        "Ratings": null,
        "Questions": null
    }
]
```

### Responses

| HTTP Status Code | Meaning                                                 | Description     | Data schema |
| ---------------- | ------------------------------------------------------- | --------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Get Dentist ALL | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name                   | Type    | Required | Restrictions | Title | description |
| ---------------------- | ------- | -------- | ------------ | ----- | ----------- |
| » ID                   | integer | true     | none         |       | none        |
| » CreatedAt            | string  | true     | none         |       | none        |
| » UpdatedAt            | string  | true     | none         |       | none        |
| » DeletedAt            | null    | true     | none         |       | none        |
| » Name                 | string  | true     | none         |       | none        |
| » Specialist           | string  | true     | none         |       | none        |
| » AboutMe              | string  | true     | none         |       | none        |
| » WorkYearExperience   | integer | true     | none         |       | none        |
| » PatientCount         | integer | true     | none         |       | none        |
| » Picture              | string  | true     | none         |       | none        |
| » OnlineConsultations  | null    | true     | none         |       | none        |
| » OfflineConsultations | null    | true     | none         |       | none        |
| » Appointments         | null    | true     | none         |       | none        |
| » Ratings              | null    | true     | none         |       | none        |
| » Questions            | null    | true     | none         |       | none        |

## GET Get Dentist by ID

GET /dentists/2

> Body Parameters

```json
{}
```

### Params

| Name | Location | Type   | Required | Description |
| ---- | -------- | ------ | -------- | ----------- |
| body | body     | object | no       | none        |

> Response Examples

> Get Dentist by ID

```json
{
    "ID": 2,
    "CreatedAt": "2024-01-12T17:26:49.131+07:00",
    "UpdatedAt": "2024-01-12T17:26:49.131+07:00",
    "DeletedAt": null,
    "Name": "Dr Luna Lovegood",
    "Specialist": "Paediatric",
    "AboutMe": "I am a dentist at Saiful Anwar Hospital in Malang, Indonesia. I have been working here for three years, and I have treated hundreds of patients. I am passionate about dentistry, and I believe that everyone deserves to have healthy, beautiful teeth.",
    "WorkYearExperience": 5,
    "PatientCount": 200,
    "Picture": "",
    "OnlineConsultations": [
        {
            "ID": 3,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Day": "Wednesday",
            "WorkHour": "09:00-15:00",
            "Price": 25000,
            "DentistID": 2,
            "ServiceID": 2,
            "Appointment": null
        },
        {
            "ID": 4,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Day": "Friday",
            "WorkHour": "09:00-15:00",
            "Price": 25000,
            "DentistID": 2,
            "ServiceID": 2,
            "Appointment": null
        },
        {
            "ID": 5,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Day": "Wednesday",
            "WorkHour": "09:00-15:00",
            "Price": 25000,
            "DentistID": 2,
            "ServiceID": 3,
            "Appointment": null
        },
        {
            "ID": 6,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Day": "Friday",
            "WorkHour": "09:00-15:00",
            "Price": 25000,
            "DentistID": 2,
            "ServiceID": 3,
            "Appointment": null
        },
        {
            "ID": 7,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Day": "Wednesday",
            "WorkHour": "09:00-15:00",
            "Price": 25000,
            "DentistID": 2,
            "ServiceID": 4,
            "Appointment": null
        },
        {
            "ID": 8,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Day": "Friday",
            "WorkHour": "09:00-15:00",
            "Price": 25000,
            "DentistID": 2,
            "ServiceID": 4,
            "Appointment": null
        }
    ],
    "OfflineConsultations": [
        {
            "ID": 1,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Place": "Dr. Saiful Anwar General Hospital",
            "Day": "Monday",
            "WorkHour": "09:00-12:00",
            "Price": 25000,
            "DentistID": 2,
            "ServiceID": 5,
            "Appointment": null
        },
        {
            "ID": 2,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Place": "Dr. Saiful Anwar General Hospital",
            "Day": "Thursday",
            "WorkHour": "09:00-12:00",
            "Price": 25000,
            "DentistID": 2,
            "ServiceID": 5,
            "Appointment": null
        },
        {
            "ID": 3,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Place": "Dr. Saiful Anwar General Hospital",
            "Day": "Tuesday",
            "WorkHour": "12:00-15:00",
            "Price": 25000,
            "DentistID": 2,
            "ServiceID": 5,
            "Appointment": null
        }
    ],
    "Appointments": null,
    "Ratings": null,
    "Questions": null
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description       | Data schema |
| ---------------- | ------------------------------------------------------- | ----------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Get Dentist by ID | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name                   | Type     | Required | Restrictions | Title | description |
| ---------------------- | -------- | -------- | ------------ | ----- | ----------- |
| » ID                   | integer  | true     | none         |       | none        |
| » CreatedAt            | string   | true     | none         |       | none        |
| » UpdatedAt            | string   | true     | none         |       | none        |
| » DeletedAt            | null     | true     | none         |       | none        |
| » Name                 | string   | true     | none         |       | none        |
| » Specialist           | string   | true     | none         |       | none        |
| » AboutMe              | string   | true     | none         |       | none        |
| » WorkYearExperience   | integer  | true     | none         |       | none        |
| » PatientCount         | integer  | true     | none         |       | none        |
| » Picture              | string   | true     | none         |       | none        |
| » OnlineConsultations  | [object] | true     | none         |       | none        |
| »» ID                  | integer  | true     | none         |       | none        |
| »» CreatedAt           | string   | true     | none         |       | none        |
| »» UpdatedAt           | string   | true     | none         |       | none        |
| »» DeletedAt           | null     | true     | none         |       | none        |
| »» Day                 | string   | true     | none         |       | none        |
| »» WorkHour            | string   | true     | none         |       | none        |
| »» Price               | integer  | true     | none         |       | none        |
| »» DentistID           | integer  | true     | none         |       | none        |
| »» ServiceID           | integer  | true     | none         |       | none        |
| »» Appointment         | null     | true     | none         |       | none        |
| » OfflineConsultations | [object] | true     | none         |       | none        |
| »» ID                  | integer  | true     | none         |       | none        |
| »» CreatedAt           | string   | true     | none         |       | none        |
| »» UpdatedAt           | string   | true     | none         |       | none        |
| »» DeletedAt           | null     | true     | none         |       | none        |
| »» Place               | string   | true     | none         |       | none        |
| »» Day                 | string   | true     | none         |       | none        |
| »» WorkHour            | string   | true     | none         |       | none        |
| »» Price               | integer  | true     | none         |       | none        |
| »» DentistID           | integer  | true     | none         |       | none        |
| »» ServiceID           | integer  | true     | none         |       | none        |
| »» Appointment         | null     | true     | none         |       | none        |
| » Appointments         | null     | true     | none         |       | none        |
| » Ratings              | null     | true     | none         |       | none        |
| » Questions            | null     | true     | none         |       | none        |

## PUT Answer Question

PUT /dentist/2/question/2/answer

> Body Parameters

```json
{
    "answer": "The symptoms of tooth decay can vary depending on the severity of the decay. In the early stages, tooth decay may not cause any symptoms at all. However, as the decay progresses, you may begin to experience the following symptoms: Sensitivity to hot or cold foods and drinks Pain when biting or chewingLoose or chipped teeth Dark spots or cavities on the teeth If you experience any of these symptoms, it is important to see a dentist as soon as possible. Tooth decay can be a serious condition that can lead to tooth loss if left untreated."
}
```

### Params

| Name     | Location | Type   | Required | Description |
| -------- | -------- | ------ | -------- | ----------- |
| body     | body     | object | no       | none        |
| » answer | body     | string | yes      | none        |

> Response Examples

> New Request

```json
{
    "message": "Question answered successfully"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | New Request | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

## PUT Answer Question

PUT /dentist/2/question/1/answer

> Body Parameters

```json
{
    "answer": "The symptoms of tooth decay can vary depending on the severity of the decay. In the early stages, tooth decay may not cause any symptoms at all. However, as the decay progresses, you may begin to experience the following symptoms: Sensitivity to hot or cold foods and drinks Pain when biting or chewingLoose or chipped teeth Dark spots or cavities on the teeth If you experience any of these symptoms, it is important to see a dentist as soon as possible. Tooth decay can be a serious condition that can lead to tooth loss if left untreated."
}
```

### Params

| Name     | Location | Type   | Required | Description |
| -------- | -------- | ------ | -------- | ----------- |
| body     | body     | object | no       | none        |
| » answer | body     | string | yes      | none        |

> Response Examples

> New Request

```json
{
    "message": "Question answered successfully"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | New Request | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

## POST Appointment

POST /profile/3/appointment

> Body Parameters

```json
{
    "OnlineConsultationID": 4
}
```

### Params

| Name                   | Location | Type    | Required | Description |
| ---------------------- | -------- | ------- | -------- | ----------- |
| body                   | body     | object  | no       | none        |
| » OnlineConsultationID | body     | integer | yes      | none        |

> Response Examples

> Appointment

```json
{
    "id": 13
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Appointment | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name | Type    | Required | Restrictions | Title | description |
| ---- | ------- | -------- | ------------ | ----- | ----------- |
| » id | integer | true     | none         |       | none        |

## PUT BAYARR

PUT /payments/1

> Body Parameters

```yaml
Method: GOPAY
photo: string
```

### Params

| Name     | Location | Type           | Required | Description |
| -------- | -------- | -------------- | -------- | ----------- |
| body     | body     | object         | no       | none        |
| » Method | body     | string         | yes      | none        |
| » photo  | body     | string(binary) | yes      | none        |

> Response Examples

> New Request

```json
{
    "payment": {
        "ID": 1,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2024-01-15T23:03:17.991+07:00",
        "DeletedAt": null,
        "Amount": 0,
        "Status": false,
        "Method": "",
        "Photo": "https://storage.googleapis.com/supple-hulling-408914.appspot.com/Screenshot%202023-04-25%20150604.png?Expires=1705420997&GoogleAccessId=supple-hulling-408914%40appspot.gserviceaccount.com&Signature=pvpZd%2Bp%2BG%2B3nLa8xZuvrQNEsiCS8DzFYmisyOM%2BVi61sIJ5S42tUGhOfOfI8bkW7r6mqTpPyDGh6R9EQqZ5fOWFFd0GEvAIP5O2d4FMxZHU5u6kNuiwOm3Ga13ScSKsTahcCsccRqPYjNL%2BhfT92isGX6UXvnhcHL2LI2JnAa6HhXdVE7tuThK%2FijooMaLFoYj3u%2FpTvyZd52mDoV63llQaTnMZ47%2BLWuhCnTbkKfKWViuYagBAP6LlxNcKhM6%2BU44M92TdBLtWjnqIF779LP6dmKyIqI3KBQK5VvZ3f%2F7nTtRz9qt%2Bw7k6w78Bv5DRNHp2xcbx9Ox7gJMrGG8Ut3A%3D%3D",
        "AppointmentID": 0
    }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | New Request | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name             | Type    | Required | Restrictions | Title | description |
| ---------------- | ------- | -------- | ------------ | ----- | ----------- |
| » payment        | object  | true     | none         |       | none        |
| »» ID            | integer | true     | none         |       | none        |
| »» CreatedAt     | string  | true     | none         |       | none        |
| »» UpdatedAt     | string  | true     | none         |       | none        |
| »» DeletedAt     | null    | true     | none         |       | none        |
| »» Amount        | integer | true     | none         |       | none        |
| »» Status        | boolean | true     | none         |       | none        |
| »» Method        | string  | true     | none         |       | none        |
| »» Photo         | string  | true     | none         |       | none        |
| »» AppointmentID | integer | true     | none         |       | none        |

## GET Get Appointment

GET /profile/2/appointment/8

> Body Parameters

```json
{}
```

### Params

| Name | Location | Type   | Required | Description |
| ---- | -------- | ------ | -------- | ----------- |
| body | body     | object | no       | none        |

> Response Examples

> Appointment Copy

```json
{
    "appointment": {
        "ID": 11,
        "CreatedAt": "2024-02-20T21:16:28.099+07:00",
        "UpdatedAt": "2024-02-20T21:16:28.099+07:00",
        "DeletedAt": null,
        "PatientName": "Rapi Kuntadi",
        "DentistID": 2,
        "OnlineConsultationID": 4,
        "OfflineConsultationID": null,
        "Day": "Friday",
        "WorkHour": "09:00-15:00",
        "TotalPrice": 25000,
        "ProfileID": 3
    }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description      | Data schema |
| ---------------- | ------------------------------------------------------- | ---------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Appointment Copy | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name                     | Type    | Required | Restrictions | Title | description |
| ------------------------ | ------- | -------- | ------------ | ----- | ----------- |
| » appointment            | object  | true     | none         |       | none        |
| »» ID                    | integer | true     | none         |       | none        |
| »» CreatedAt             | string  | true     | none         |       | none        |
| »» UpdatedAt             | string  | true     | none         |       | none        |
| »» DeletedAt             | null    | true     | none         |       | none        |
| »» PatientName           | string  | true     | none         |       | none        |
| »» DentistID             | integer | true     | none         |       | none        |
| »» OnlineConsultationID  | integer | true     | none         |       | none        |
| »» OfflineConsultationID | null    | true     | none         |       | none        |
| »» Day                   | string  | true     | none         |       | none        |
| »» WorkHour              | string  | true     | none         |       | none        |
| »» TotalPrice            | integer | true     | none         |       | none        |
| »» ProfileID             | integer | true     | none         |       | none        |

## GET Get Appointment All

GET /profile/1/appointments

> Body Parameters

```json
{}
```

### Params

| Name | Location | Type   | Required | Description |
| ---- | -------- | ------ | -------- | ----------- |
| body | body     | object | no       | none        |

> Response Examples

> Get Appointment All

```json
{
    "appointments": [
        {
            "ID": 1,
            "CreatedAt": "2024-03-01T01:18:24.494Z",
            "UpdatedAt": "2024-03-01T01:18:24.494Z",
            "DeletedAt": null,
            "PatientName": "RAVI KUNTADI",
            "DentistID": 1,
            "OnlineConsultationID": 1,
            "OfflineConsultationID": null,
            "Day": "Tuesday",
            "WorkHour": "10:00-18:00",
            "TotalPrice": 80,
            "ProfileID": 1
        }
    ]
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description         | Data schema |
| ---------------- | ------------------------------------------------------- | ------------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Get Appointment All | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name                     | Type     | Required | Restrictions | Title | description |
| ------------------------ | -------- | -------- | ------------ | ----- | ----------- |
| » appointments           | [object] | true     | none         |       | none        |
| »» ID                    | integer  | false    | none         |       | none        |
| »» CreatedAt             | string   | false    | none         |       | none        |
| »» UpdatedAt             | string   | false    | none         |       | none        |
| »» DeletedAt             | null     | false    | none         |       | none        |
| »» PatientName           | string   | false    | none         |       | none        |
| »» DentistID             | integer  | false    | none         |       | none        |
| »» OnlineConsultationID  | integer  | false    | none         |       | none        |
| »» OfflineConsultationID | null     | false    | none         |       | none        |
| »» Day                   | string   | false    | none         |       | none        |
| »» WorkHour              | string   | false    | none         |       | none        |
| »» TotalPrice            | integer  | false    | none         |       | none        |
| »» ProfileID             | integer  | false    | none         |       | none        |

## PUT BAYARR

PUT /payments/2

> Body Parameters

```yaml
Method: GOPAY
photo: string
```

### Params

| Name     | Location | Type           | Required | Description |
| -------- | -------- | -------------- | -------- | ----------- |
| body     | body     | object         | no       | none        |
| » Method | body     | string         | no       | none        |
| » photo  | body     | string(binary) | no       | none        |

> Response Examples

> BAYARR

```json
{
    "payment": {
        "ID": 2,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2024-02-19T20:11:02.161+07:00",
        "DeletedAt": null,
        "Amount": 0,
        "Status": true,
        "Method": "GOPAY",
        "Photo": "https://storage.googleapis.com/supple-hulling-408914.appspot.com/KT%7B.jpg?Expires=1708434662&GoogleAccessId=supple-hulling-408914%40appspot.gserviceaccount.com&Signature=UO4iKjoeqIXq4gGK9OH5CG5SV7YTN4nlNCj%2BuF4cEpGAwf0A82HezeWeLBnDqduUHgHzgFudIZ%2BuNP2zOyRTWk4uzI35DgxNZifmGymsOOesfIjckG5YHHxZf5pHvPW%2FPxGki0KN8oSwz%2B0jzjuz2oAvn6kzY2qoJoE8FGmkXt5jlgt8laCtg29lfdzo07o6k%2BXaDrmLSk5ctYwT51CXm31ITmsY7uE%2FhGtf%2B%2BoArjLLrTWDA2RonJ1wf3NkVHzFSLA8wglCveeBfEEhPtJ1VO8PZ0JEFXSwtbtzvXgbo6JGgz6PFIsdpBCGsiZd6p0yuZ%2FMIo7CgUEOSSEijRJ1%2Bg%3D%3D",
        "AppointmentID": 0
    }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | BAYARR      | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name             | Type    | Required | Restrictions | Title | description |
| ---------------- | ------- | -------- | ------------ | ----- | ----------- |
| » payment        | object  | true     | none         |       | none        |
| »» ID            | integer | true     | none         |       | none        |
| »» CreatedAt     | string  | true     | none         |       | none        |
| »» UpdatedAt     | string  | true     | none         |       | none        |
| »» DeletedAt     | null    | true     | none         |       | none        |
| »» Amount        | integer | true     | none         |       | none        |
| »» Status        | boolean | true     | none         |       | none        |
| »» Method        | string  | true     | none         |       | none        |
| »» Photo         | string  | true     | none         |       | none        |
| »» AppointmentID | integer | true     | none         |       | none        |

## PUT BAYARR Dummy

PUT /paymentsdummy/3

> Body Parameters

```yaml
Method: GOPAY
photo: string
```

### Params

| Name     | Location | Type           | Required | Description |
| -------- | -------- | -------------- | -------- | ----------- |
| body     | body     | object         | no       | none        |
| » Method | body     | string         | yes      | none        |
| » photo  | body     | string(binary) | yes      | none        |

> Response Examples

> BAYARR Dummy

```json
{
    "message": "payment successfully"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description  | Data schema |
| ---------------- | ------------------------------------------------------- | ------------ | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | BAYARR Dummy | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

## POST CARIES DETECTOR

POST /upload

> Body Parameters

```yaml
myFile: string
text: I have had toothache in my right rear molar since yesterday. The pain
    feels like aching and throbbing, especially when I eat or drink something cold
    or hot. The pain also spread to my ears and jaw. When I checked my teeth, I
    found that there was a small hole in the affected part of the tooth. The hole
    measures about 1 millimeter.
```

### Params

| Name     | Location | Type           | Required | Description |
| -------- | -------- | -------------- | -------- | ----------- |
| body     | body     | object         | no       | none        |
| » myFile | body     | string(binary) | yes      | none        |
| » text   | body     | string         | yes      | none        |

> Response Examples

> CARIES DETECTOR

```json
{
    "persen": "25% ",
    "text": "\nThere are a few things you can do to treat your toothache and prevent further damage to your tooth. \nFirst, you should see a dentist as soon as possible. They will be able to examine your tooth and determine the best course of treatment. In the meantime \nYou can take over-the-counter pain relievers, such as ibuprofen or ace \nIt is important to note that these are only temporary measures. You should still see a dentist as soon as possible to get the best possible treatment for your toothache. "
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description     | Data schema |
| ---------------- | ------------------------------------------------------- | --------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | CARIES DETECTOR | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name     | Type   | Required | Restrictions | Title | description |
| -------- | ------ | -------- | ------------ | ----- | ----------- |
| » persen | string | true     | none         |       | none        |
| » text   | string | true     | none         |       | none        |

## GET New Request Copy

GET /

> Response Examples

> 200 Response

```json
{}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

# Data Schema

<h2 id="tocS_Tag">Tag</h2>

<a id="schematag"></a>
<a id="schema_Tag"></a>
<a id="tocStag"></a>
<a id="tocstag"></a>

```json
{
    "id": 1,
    "name": "string"
}
```

### Attribute

| Name | Type           | Required | Restrictions | Title | Description |
| ---- | -------------- | -------- | ------------ | ----- | ----------- |
| id   | integer(int64) | false    | none         |       | Tag ID      |
| name | string         | false    | none         |       | Tag Name    |

<h2 id="tocS_Category">Category</h2>

<a id="schemacategory"></a>
<a id="schema_Category"></a>
<a id="tocScategory"></a>
<a id="tocscategory"></a>

```json
{
    "id": 1,
    "name": "string"
}
```

### Attribute

| Name | Type           | Required | Restrictions | Title | Description   |
| ---- | -------------- | -------- | ------------ | ----- | ------------- |
| id   | integer(int64) | false    | none         |       | Category ID   |
| name | string         | false    | none         |       | Category Name |

<h2 id="tocS_Pet">Pet</h2>

<a id="schemapet"></a>
<a id="schema_Pet"></a>
<a id="tocSpet"></a>
<a id="tocspet"></a>

```json
{
    "id": 1,
    "category": {
        "id": 1,
        "name": "string"
    },
    "name": "doggie",
    "photoUrls": ["string"],
    "tags": [
        {
            "id": 1,
            "name": "string"
        }
    ],
    "status": "available"
}
```

### Attribute

| Name      | Type                        | Required | Restrictions | Title | Description      |
| --------- | --------------------------- | -------- | ------------ | ----- | ---------------- |
| id        | integer(int64)              | true     | none         |       | Pet ID           |
| category  | [Category](#schemacategory) | true     | none         |       | group            |
| name      | string                      | true     | none         |       | name             |
| photoUrls | [string]                    | true     | none         |       | image URL        |
| tags      | [[Tag](#schematag)]         | true     | none         |       | tag              |
| status    | string                      | true     | none         |       | Pet Sales Status |

#### Enum

| Name   | Value     |
| ------ | --------- |
| status | available |
| status | pending   |
| status | sold      |
