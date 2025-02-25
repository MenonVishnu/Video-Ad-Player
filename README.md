# Video-Ad-Player

A Fullstack application which plays a Video and displays ad-overlay fetched from backend across random position. The application also tracks when the user clicks on the advertisement and sends meta data to backend.

## Table of Content

- Run Locally

- Run using Docker

- API Documentation

- Architecture

## Run Locally

To deploy this project locally follow the below steps:

Open terminal where your project is downloaded.

Go to Backend folder

```bash
  cd backend
```

Install go packages

```bash
  go mod download
```

Start backend server at port 8080

```bash
  go run main.go
```

This should start the backend server at http://localhost:8080/

Go to Frontend folder

```bash
  cd ../frontend
```

Install node_modules

```bash
  npm install
```

Start React server at port 3000

```bash
  npm start
```

This should start your React server at http://localhost:3000/

## Run using Docker

To deploy this project using Docker follow the below steps:

Build the docker image

```bash
  docker compose up --build
```

Once build is complete, your servers would start and you can access the application through: http://localhost:3000

## API Documentation

#### Get advertisement

Retrieves a list of advertisements including their image URLs and target URLs.

```http
  GET /api/v1/ads
```

Request:

- Method: `GET`

- Endpoint: `https://localhost:8080/api/v1/ads`

Response:

- Status Code: `200 OK` (On success)

- Content-Type: `application/json`

Response Body:

```
{
    "message": "<message>",
    "data": [
        {
            "ad_id": int,
            "image_url": string,
            "target_url": string
        },
        {
             "ad_id": int,
            "image_url": string,
            "target_url": string
        }
    ],
    "error": "<error_message>"
}
```

Error Responses:

`500` – If there is a server error.

#### POST user clicks

Logs when a user clicks on an advertisement, capturing details like the ad ID, timestamp, IP address, and video timestamp.

```http
  POST /api/v1/ads/click
```

Request:

- Method: `POST`

- Endpoint: `https://localhost:8080/api/v1/ads/click`

- Content-Type: `application/json`

Request Body

```
{
    "ad_id": int,
    "timestamp": string,
    "ip": string,
    "video_timestamp": float
}
```

Response:

- Status Code: `200 OK` (On success)

- Content-Type: `application/json`

Response Body:

```
{
    "message": string,
    "data": interface{},
    "error": interface{}
}
```

Error Responses:

`400` - If required fields are missing.

`500` – If there is a server error.

## Architecture

![Architecture ](https://drive.google.com/file/d/1Mv-NitjVEG14ntSGiTbXKuXnLHC3X68A/view?usp=drive_link)

#### Frontend

- Contains a Video Player.
- Initially calls an API to fetch advertisements from the backend.
- Once ads are fetched, the video player starts, and ad overlays are shown in random positions on the window.
- When a user clicks on an ad, an API call logs the following data:

```
{
    ad_id
    timestamp
    ip
    video_timestamp
}
```

#### Backend

- On Startup:
  - Checks if the advertisement and clickdata tables exist, creating them if necessary.
  - Loads dummy advertisements from dummydata.json into the advertisements table if not already present.
- GET /ads: Fetches all advertisements from the advertisement table and returns them as a response.
- POST /ads/click: Stores {ad_id, timestamp, video_timestamp} from the request body, retrieves the IP from the request headers, and stores it in the clickdata table.
