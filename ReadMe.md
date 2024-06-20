# Gullit-Pen

This repository contains a simple Go web server (`Main.go`) that serves various HTML files located in the `./static` directory based on different HTTP routes. These HTML files are built using Tailwind CSS and are stored in the `./static` directory.

## Description

The `Main.go` file sets up a Go HTTP server that listens on port `8080` and handles several routes:

- `/GullitPen` serves `./static/seven.html`
- `/seven20` serves `./static/seven20.html`
- `/seven2` serves `./static/seven2.html`
- `/seven` serves `./static/seven.html`
- `/seven0` serves `./static/seven0.html`
- `/seven1` serves `./static/seven1.html`
- `/seven10` serves `./static/seven10.html`

Additionally, any other route will serve static files from the `./static` directory using `http.FileServer`. These HTML files in the `./static` directory are designed using Tailwind CSS.

## Usage

1. Ensure you have Go installed on your machine.
2. Clone this repository:
   ```bash
   git clone https://github.com/samyam81/Gullit-Pen
   cd Gullit-Pen
   ```
3. Run the server:
   ```bash
   go run Main.go
   ```
4. Open your web browser and navigate to `http://localhost:8080/GullitPen` to see the server in action.

## Screenshots

### Dark Mode

![Seven0 Dark Mode](https://github.com/samyam81/Gullit-Pen/raw/main/output/Dark/Seven0.png)
![Seven10 Dark Mode](https://github.com/samyam81/Gullit-Pen/raw/main/output/Dark/Seven10.png)
![Seven20 Dark Mode](https://github.com/samyam81/Gullit-Pen/raw/main/output/Dark/Seven20.png)

### Light Mode

![Seven Light Mode](https://github.com/samyam81/Gullit-Pen/raw/main/output/Light/Seven.png)
![Seven1 Light Mode](https://github.com/samyam81/Gullit-Pen/raw/main/output/Light/Seven1.png)
![Seven2 Light Mode](https://github.com/samyam81/Gullit-Pen/raw/main/output/Light/Seven2.png)

## Author

This repository is maintained by [Samyam](https://github.com/samyam81).

## Tailwind CSS

The HTML files in the `./static` directory are styled using Tailwind CSS. Tailwind CSS is a utility-first CSS framework for building custom designs without writing any CSS.

To learn more about Tailwind CSS, visit [tailwindcss.com](https://tailwindcss.com/).
