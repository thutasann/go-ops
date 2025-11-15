function immutable() {
  const user = { name: 'Alice', age: 25 };
  const updatedUser = { ...user, age: 26 };

  console.log(user);
  console.log(updatedUser);
}

function mutable() {
  const user = { name: 'Alice', age: 25 };

  user.age = 26;

  console.log(user);
}

function immutable_updates_with_arrays() {
  const nums = [1, 2, 3];
  const newNums = [...nums, 4];
  console.log(nums);
  console.log(newNums);
}

// immutable_updates_with_arrays();
