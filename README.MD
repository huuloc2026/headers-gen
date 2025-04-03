# hdgen - Comment Block Header Generator

`hdgen` is a command-line tool that generates structured headers for source code files in multiple programming languages, including Go, Java, Python, Rust, and TypeScript.

## 🚀 Installation

### 1️⃣ Install from source

```sh
git clone https://github.com/huuloc2026/headers-gen.git
cd hd-gen
go build -o hdgen ./cmd
```
This will generate the hdgen executable in your current directory.

## 🔧 Usage

After successfully installing the hd-gen package, you can use it to generate structured headers for different programming languages.
Basic Command Structure

To use hd-gen, use the following structure:
```
hd-gen -<language> "<header-title>"
```
Replace <language> with one of the supported programming languages and provide a title for the header.
Supported Languages

- go - for Go language

- java - for Java language

- python - for Python language

- rust - for Rust language

- ts - for TypeScript language


# Examples:

## Go:

```
hd-gen -go "Variable Declaration"
```
- Output:
```
//************************************************************
//*                    VARIABLE DECLARATION                  *
//************************************************************
```

## Java:
```
hd-gen -java "JAKE ONYX"
```
- Output:
```
//************************************************************
//*                        JAKE ONYX                         *
//************************************************************
```
## - Python:

```
hd-gen -python "JAKE ONYX"
```
- Output:
```
#############################################################
#                         JAKE ONYX                         #
#############################################################
```

## -  Rust:
```
hd-gen -rust "JAKE ONYX"
```
- Output:
```
//************************************************************
//*                        JAKE ONYX                         *
//************************************************************
```
## - TypeScript:
```
hd-gen -ts "JAKE ONYX"
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

