console.log('Start');

// Asynchronous task
setTimeout(() => {
  console.log('Inside setTimeout (after 2 sec)');
}, 2000);

console.log('End');

async function fetchData() {
  console.log('Start');

  const data = await new Promise((resolve) => setTimeout(() => resolve('Fetched async data'), 1500));

  console.log(data);
  console.log('End');
}

fetchData();
