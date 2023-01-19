<template>
  <div class="gradient"></div>
</template>

<script>
export default {
  name: 'GradientTest',
};
</script>

<style scoped>
.gradient {
  background-color: rgb(252, 252, 252);
  z-index: -1000;
  display: flex;
  place-content: center;
  align-items: center;
  --size: 2000px;
  --speed: 50s;
  --easing: cubic-bezier(0.8, 0.2, 0.2, 0.8);
  width: var(--size);
  height: var(--size);
  filter: blur(calc(var(--size) / 5));
  background-image: linear-gradient(#37eba9, #5b37eb);
  animation: rotate var(--speed) var(--easing) alternate infinite;
  border-radius: 30% 70% 70% 30% / 30% 30% 70% 70%;
}
@media (min-width: 720px) {
  .gradient {
    --size: 2000px;
  }
}
@keyframes rotate {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
