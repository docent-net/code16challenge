# code16challenge

Porozmawiajmy o Linuksie #pdk, a dokładniej to o kontenerach linuksowych. I
zróbmy to w ramach wspierania ciężko chorej Majki ❤

Zabawa polega na tym, że ktoś nas wskazuje, my piszemy 16 linijek dowolnego
kodu i wspieramy Maję:

- Szczegóły akcji: [https://www.siepomaga.pl/code16challenge]
- Szczegóły zabawy: [https://www.code16challenge.pl]

Przedstawiam Wam bieda wersję kontenerów linuksowych (aka Dockera). Potrafi w
kilka namespace'ów: mount, PID oraz UTS. Czyli, że kontener ma własnego
hostname'a, własnego chroota oraz własny widok na procesy. Dałoby się więcej
(jeszcze cgroupy aby były swoje zasoby i może net aby sieć była), ale brakło mi
linijek. To chyba lepsze niż białe znaki w YAMLu 😛 Szczerze dziękuję, bo miałem
powód, aby zejść dość nisko, przypomnieć sobie jak się robi chroota i jak się
forkuje i poskładać to do kupy w Golangu. Jest to na tyle ciekawy temat, że
chyba nakręcę z tego patostreama na YT gdzie pokażę jak to działa 😉

# howto

1. Potrzebujesz golanga do uruchomienia main.go
1. Potrzebujesz katalogu "cfs", do którego skopiujesz kilka binarek (sh,
hostname, ps, ls i co tam jeszcze chcesz odpalić w kontenerze) oraz bibliotek,
z których te binarki korzystają (możesz sprawdzić, które są wymagane na kilka
sposobów - najprościej będzie z ldd, np. `ldd /bin/sh`)
1. Odpalanie z roota (bo nie miałem miejsca w 16 linijkach na rootless): `go run main.go`