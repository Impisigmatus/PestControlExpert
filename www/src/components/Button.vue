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
/* .button {
  background-color: transparent;
  color: rgb(250, 250, 250);
  border: transparent;
  text-transform: uppercase;
} */
.button {
  background-color: transparent;
  border: 2px solid rgb(255, 255, 255);
  border-radius: 0.6em;
  color: rgb(255, 255, 255);
  display: flex;
  align-self: center;
  font-size: 12px;
  line-height: 1;
  padding: 10px 15px;
  text-decoration: none;
  text-align: center;
  text-transform: uppercase;
  font-weight: 700;
  transition: box-shadow 300ms ease-in-out, color 300ms ease-in-out;
}

.button:hover {
  box-shadow: 0 0 40px 40px rgb(255, 254, 254) inset;
}
.button:hover,
.button:focus {
  color: #b1c8f3;
  outline: 0;
}
</style>
