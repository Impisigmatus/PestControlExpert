<template>
  <button class="button">
    <slot></slot>
  </button>
</template>

<script>
export default {
  name: 'PceButton',
};
</script>

<style scoped>
.button {
  background-color: transparent;
  color: rgb(41, 44, 255);
  border: transparent;
  text-transform: uppercase;
}
</style>
