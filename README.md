# map-command-api
the command part of CQRS, takes in any protobuf(as bytes) and wraps a [command proto](https://github.com/Mattilsynet/mapis/blob/master/apis/command/v1/command.proto) around it. And further publishes it with jetstream.Publish(needs to be a interest based workqueue) to the kind defined in CommandSpec.type. i.e., ManagedEnvironment will be sent to map.ManagedEnvironment.<Operation> (apply or delete).
Check out wadm.yaml for further info.
## requirements

tinygo v0.33  
wit-deps(optional) : https://github.com/bytecodealliance/wit-deps  
wit-bindgen-go: https://github.com/bytecodealliance/wasm-tools-go/tree/main (clone and go install from cmd/)  

## to make it work

1. wash build  
2. wash up (in other terminal)  
3. wash app deploy wadm.yaml  
