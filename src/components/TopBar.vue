<template>
  <div class="navbar">
    <img
      class="navbar__img"
      src="@/assets/logo.png"
      alt="#"
      @click="$router.push('/')"
    />
    <!-- <div class="logo"></div> -->
    <div class="navbar__btns">
      <pce-button class="btn__nav" @click="$router.push('/price')"
        >Цены</pce-button
      >
      <pce-button
        class="btn__nav"
        style="margin-left: 20px"
        @click="$router.push('/map')"
        >Карта покрытия</pce-button
      >
    </div>
  </div>
</template>

<script>
export default {
  name: 'PceTopBar',
};
</script>

<style scoped>
.navbar {
  height: 100px;
  /* box-shadow: 1px 2px 4px rgb(120, 188, 252); */
  display: flex;
  padding: 15px 30px 15px 30px;
  flex-direction: column;
  align-items: center;
  padding: 0 15px;
  background-color: #fff;
}

/* .logo {
  margin-top: 12px;
  text-transform: uppercase;
  color: rgb(50, 81, 255);
  font-size: 22px;
  font-weight: 700;
} */

.navbar__img {
  width: 70px;
  height: 60px;
}
</style>
