gcc -std=c99 -Wall -Wextra `
    -nostdlib -ffreestanding -mconsole -Os `
    -fno-stack-check -fno-stack-protector -mno-stack-arg-probe `
    -o bare.exe main.c `
    -lkernel32