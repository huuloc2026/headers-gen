# hdgen - Comment Header Generator

`hdgen` is a command-line tool that generates structured headers for source code files in multiple programming languages, including Go, Java, Python, Rust, and TypeScript.


## 🚀 Installation

### 1️⃣ Install from source

```sh
git clone https://github.com/huuloc2026/headers-gen.git
cd headers-gen
go build -o hdgen ./cmd
```
This will generate the hdgen executable in your current directory.
#### Notes: If you want to use this CLI globally:
## For Windows:
- Move hdgen.exe to the system directory:
```sh
move hdgen.exe C:\Windows\System32\
```
## Linux/MacOs
- Step 1: move hdgen to /usr/local/bin/:
```sh
sudo mv hdgen /usr/local/bin/
```
- Step 2: Check if it's working:
```sh
hdgen -go "Have a good day"
```
## 🔧 Usage

After successfully installing the hdgen package, you can use it to generate structured headers for different programming languages.
Basic Command Structure

To use hdgen, use the following structure:
```sh
hdgen -<language> "<header-title>"
```
Replace <language> with one of the supported programming languages and provide a title for the header.
Supported Languages

- `go` - for Go language

- `java` - for Java language

- `python` - for Python language

- `rust` - for Rust language

- `ts` - for TypeScript language


# Examples:

## Go:

```sh
hdgen -go "Variable Declaration"
```
- Output:
```
//************************************************************
//*                    VARIABLE DECLARATION                  *
//************************************************************
```

## Java:
```sh
hdgen -java "JAKE ONYX"
```
- Output:
```
//************************************************************
//*                        JAKE ONYX                         *
//************************************************************
```
## Python:

```sh
hdgen -python "JAKE ONYX"
```
- Output:
```
#############################################################
#                         JAKE ONYX                         #
#############################################################
```

## Rust:
```sh
hdgen -rust "JAKE ONYX"
```
- Output:
```
//************************************************************
//*                        JAKE ONYX                         *
//************************************************************
```
## TypeScript:
```sh
hdgen -ts "JAKE ONYX"
```
- Output:
```
//************************************************************
//*                        JAKE ONYX                         *
//************************************************************
```

## 🎨 Customization

The script automatically converts the header title to uppercase, regardless of the case you provide. You can specify the title in any case (lowercase, mixed case, or uppercase), and the header will always display it in uppercase.
## Additional Features:

- Title Formatting: The header title is always centered.

- Border Style: The border style changes based on the language:

    - ```*``` for Go, Java, Rust, and TypeScript

    - ```#``` for Python
If you have any questions or issues, feel free to open an issue on the GitHub repository.
## 🤝 License

hdgen is open-source and available under the MIT License. 

This `README.md` file provides:

1. **Installation** instructions (both from source and via Go).
2. **Usage** with example commands for generating headers in different languages.
3. **Customization** and details about the header formatting.
4. **Development** instructions for contributing or modifying the code.
5. **License** information and a link to report issues.

Let me know if you need any further adjustments or additions!


### 📞 Contact:
- **Email**: [huuloc2026@gmail.com](mailto:huuloc2026@gmail.com)
- **GitHub**: [https://github.com/huuloc2026](https://github.com/huuloc2026)
- **X (Twitter)**: [x.com/huuloc2026](https://x.com/huuloc2026)
- **LinkedIn**: [https://linkedin.com/in/huuloc2026](https://linkedin.com/in/huuloc2026)

