from random import randint as r

number = r(1,10)
try:
    input = input("Guess the number 1-10: ")

    try:
        intinput = int(input)

        if intinput == number:
           print(f"Correct! The number was {number}.")
        else:
           print(f"Incorrrect, the number was {number}.")
    except ValueError:
        print("You must guess a valid number!")
except KeyboardInterrupt:
   print("\nare you trying to quit the game u stupid fuck fine but never come back")
