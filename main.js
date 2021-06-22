function isPrime(num) {
    if (num === 2) {
        return true;
    } else if (num === 1 || num % 2 === 0) {
        return false;
    } else {
        const to = Math.sqrt(num);
        for (let div = 3; div <= to; div += 2) {
            if (num % div === 0) {
                return false;
            }
        }
        return true;
    }
}

function do1(N) {
    for (let i = 0; i < N; i++) {
        const prime = isPrime(i);
        if (prime) {
            // console.log(i + ': ' + prime);
        }
    }
}

const st = Date.now();
do1(100_000_000);
console.log((Date.now() - st) / 1000);
