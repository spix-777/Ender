# Ender - AES-256 File Encryption/Decryption Tool

Ender is a command-line tool that provides AES-256 encryption and decryption for files. It uses the Cipher Block Chaining (CBC) mode of operation for encryption and the Counter (CTR) mode for decryption.

## Features

- AES-256 encryption and decryption
- Command-line interface for easy usage
- Secure random key and IV generation
- Supports both encryption and decryption operations

## Usage

To use Ender, you need to have Go installed on your machine.

1. Clone the repository and navigate to the project directory.
2. Build the executable with the following command:
    go build -o ender
3. Run the executable with appropriate flags.

### Flags

- `-v`: Print version information.
- `-e`: Encrypt a file using AES-256.
- `-d`: Decrypt a previously encrypted file.
- `-f <file_path>`: Specify the path of the target file for encryption or decryption.

### Examples

1. Encrypt a file:
    ./ender -e -f path/to/your/file.txt

2. Decrypt a file:
    ./ender -d -f path/to/encrypted/file.enc

3. Print version:
    ./ender -v


## Security Note

Make sure to keep your encryption key and IV secure. Loss of these values will result in irreversible data loss for encrypted files. Always keep a secure backup of your key and IV.

## Author

- **Author**: @spix-777
- **Version**: 1.0.0

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Have questions, suggestions, or issues? Feel free to open an issue on the [GitHub repository](https://github.com/yourusername/ender)!

