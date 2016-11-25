![Travis CI](https://travis-ci.org/travisjeffery/ec2-metadata.svg?branch=master)

## ec2-metadata

Get metadata about your running instance. 

Even works when ran inside a Docker container.

- Instance ID

    ```go
    instanceID, err := ec2metadata.InstanceID()
    # => "i-4c3dc25a", nil
    ```

- Local IP address

    ```go
    localIP, err := ec2metadata.LocalIPAddress()
    # => "10.0.0.1", nil
    ```

- Public IP address

    ```go
    publicIP, err := ec2metadata.PublicIPAddress()
    # => "72.0.0.0", nil
    ```

## License

MIT

--- 

- [travisjeffery.com](http://travisjeffery.com)
- GitHub [@travisjeffery](https://github.com/travisjeffery)
- Twitter [@travisjeffery](https://twitter.com/travisjeffery)
- Medium [@travisjeffery](https://medium.com/@travisjeffery)


