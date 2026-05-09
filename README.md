Markdown
# ASCII Art Web

## Description
**ASCII Art Web** is a web-based interface for the ASCII-art generator. This application allows users to input text, select a specific banner style (**Standard**, **Shadow**, or **Thinkertoy**), and render the output directly in their browser. The project demonstrates the implementation of a Go HTTP server using only the standard library, template rendering, and robust error handling.

## Authors
* Paul Owuor

## Usage: How to Run
1.  **Clone the project** and navigate to the root directory.
2.  **Initialize the module** (if not already done):
    ```bash
    go mod init ascii-art-web
    
Run the server:

Bash
go run main.go

4.  **Access the GUI**:
    Open your browser and navigate to `http://localhost:8080`.

## Implementation Details: Algorithm
The application follows a modular architecture to separate web logic from the ASCII generation engine:

1.  **Request Handling**: The server uses `net/http` to listen for requests. `GET /` serves the main page, while `POST /ascii-art` processes the form data.
2.  **Banner Parsing**: The logic in `internal/ascii/parser.go` reads the selected banner file from the `banners/` directory. 
3.  **ASCII Generation**: The `internal/ascii/printer.go` calculates the exact lines to pull from the banner files using the formula:  
    `LineIndex = (Character_ASCII_Value - 32) * 9 + 1`.
    It then assembles these lines horizontally to form the final art.
4.  **Template Rendering**: The result is passed to `templates/index.html`, where it is displayed inside `<pre>` tags to preserve the fixed-width formatting.

## HTTP Status Codes
*   **200 OK**: For successful requests and rendering.
*   **400 Bad Request**: For malformed requests or empty inputs.
*   **404 Not Found**: For invalid URLs or missing banner/template files.
*   **500 Internal Server Error**: For unhandled server-side issues.

## Project Structure
```text
.
├── main.go               # Server initialization and HTTP handlers
├── banners/              # Banner source files (.txt)
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── internal/
│   └── ascii/            # Business logic
│       ├── parser.go     # File reading logic
│       └── printer.go    # ASCII assembly logic
├── templates/
│   └── index.html        # HTML GUI
└── README.md
Instructions
Written entirely in Go using only standard packages.

HTML templates are stored in the /templates directory.

The code follows standard Go formatting and best practices.
