# BoxSize Service
(Docker - Makefile - golang)

BoxSize is a golang micro-service to compute the volume contains inside a box.

It accepts tree integers as parameters : _height_, _width_, _length_.

Then the service returns two strings with two informations :
The box size in cm³ and the corresponding box capacity in liters.

If 42 is a number somewhere, we raise an error because it's a special number.


## Installation

You can build and run the service inside a Docker. The makefile has a _docker-build_ and a _docker-run_ commands.
```
$> make docker-build
$> make docker-run
+ exec app
BoxSize Server will run on port 3000
```

If you prefer avoid docker, feel free to run :
```
make run
```

## Standard use-cases


### Normal case

Query :
```sh
curl -XPOST -d '{"height": 1, "width": 3, "length": 2}' localhost:3000
```
Answer :
```json
{"results":["The box volume is 6 cm³","It's also 0 liters."]}
```


### Error case <value 42>

Query :
```sh
curl -XPOST -d '{"height": 1, "width": 42, "length": 2}' localhost:3000
```

Answer :
```
The box mensuration can't have a 42 number
```



### Error case with bad keys

Query :
```sh
curl -XPOST -d '{"bad_stuff": 123}' localhost:3000
```


Answer :
```
An error occured while decoding the json format request.
```
