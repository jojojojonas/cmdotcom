# cmdotcom
With this small library we wanted to build a fast way to send SMS, WhatsApp or Viber over the interface of cm.com.

## Install
First you have to install the package:

```console
go get github.com/jojojojonas/cmdotcom
```

## How to use?
As already mentioned, you can send SMS relatively easy and get a reply about all data, therefore everything was built in structs.

You only need the token, the sender, the recipient, the message and the allowed channel.

```go
message, err := cmdotcom.NewMessage(cmdotcom.Config{"Message", []cmdotcom.To{{"Number1"}}, "Sender", "SMS", "Token"})
if err != nil {
fmt.Println(err)
}
```

## Help
If you have any questions or comments, please contact us by e-mail at [info@jj-ideenschmiede.de](mailto:info@jj-ideenschmiede.de)