
supervar = "global"

def main():
    print("START")

    supervar = "local"

    if True:
        supervar = "block"
        print(f"supervar inside block  = {supervar}")


    print(f"supervar inside main = {supervar}")

    print("END")

main()
