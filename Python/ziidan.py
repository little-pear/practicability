christopher = {}
christopher['first'] = 'Christopher'
christopher['last'] = 'Harrison'

susan = {'first' : 'Susan', 'last' : 'Ibach'}

# print(christopher)
# print(susan)

print()

people = [christopher, susan]
people.append({'first' : 'Bill', 'last' : 'Gates'})

print(people)

print()
presenters = people[0:2]

print(presenters)
print(len(presenters))
