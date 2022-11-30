<template>
  <div class="map">
    <iframe
      src="https://yandex.ru/map-widget/v1/?um=constructor%3Ae3928b383dca618308743ed387767940121ffab42b25f44f04dd31219e4cf45a&amp;source=constructor"
      width="500"
      height="372"
      frameborder="0"
    ></iframe>
  </div>
</template>

<script>
export default {
  name: 'MapPage',
};
</script>

<style scoped>
.map {
  margin-top: 100px;
}
</style>
