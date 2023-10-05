# Email Domain Checker

The Email Domain Checker is a simple Go application that checks various DNS records for a given email domain to determine if it has certain configurations like MX records, SPF (Sender Policy Framework) records, and DMARC (Domain-based Message Authentication, Reporting, and Conformance) records.

## Usage

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/Cyberguru1/Go-Email-Checker.git
   ```

```

2. Navigate to the project directory:

   ```bash
   cd Go-Email--Checker
```

3. Run the application:

   ```bash
   go run main.go
   ```
4. Enter an email domain when prompted. The application will fetch and display information about the domain's DNS records.

## Features

- Check if the provided email domain has MX (Mail Exchange) records.
- Check if the provided email domain has SPF (Sender Policy Framework) records and display the SPF record if present.
- Check if the provided email domain has DMARC (Domain-based Message Authentication, Reporting, and Conformance) records and display the DMARC record if present.

## Example

Here's an example of running the Email Domain Checker:

```bash
------------Welcome To Email Domain Checker-------------------

Enter Email Domain: example.com
```

The application will then display the results:

```


------------------ Checker Result ---------------------

Domain: example.com 
HasMx: true 
hasSPF: true 
spfRecord: v=spf1 include:_spf.google.com ~all 
hasDMARC: true 
dmarcRecord: v=DMARC1; p=none; rua=mailto:postmaster@example.com

Enter Email Domain: 
```

More Examples:

![](https://imgur.com/qgUqe4c.png)

![](https://imgur.com/RpPaMYc.png)


## Contributing

Feel free to contribute to this project by creating issues or submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

You can copy and paste this Markdown content into a README.md file in your project repository. Make sure to replace the placeholders such as the repository URL, your username, and any additional information specific to your project.
