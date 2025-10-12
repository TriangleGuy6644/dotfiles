import random

number = random.randint(1, 10)

input = input("Guess the number 1-10: ")

try:
    intinput = int(input)

    if intinput == number:
       print(f"Correct! The number was {number}.")
    else:
       print(f"Incorrrect, the number was {number}.")
except ValueError:
   print("You must guess a valid number!")