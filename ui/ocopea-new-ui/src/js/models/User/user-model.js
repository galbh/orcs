export class User {
  id;
  name;
  firstName;
  lastName;
  email;

  constructor(user) {
    this.id =        user.id;
    this.name =      user.name;
    this.firstName = user.firstName;
    this.lastName =  user.lastName;
    this.email =     user.email;
    this.fullName =  `${user.firstName} ${user.lastName}`;
  }
}
