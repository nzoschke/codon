export const Counter = () => {
  let count = $state(0);

  const inc = () => {
    count++;
  };

  return {
    get count() {
      return count;
    },
    inc,
  };
};
