<template>
  <div class="navbar">
    <img
      class="navbar__img"
      src="@/assets/logo.png"
      alt="#"
      @click="$router.push('/')"
    />
    <div class="logo">
      <h1>Pest Control Expert</h1>
    </div>
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
  background-color: #8bc6ec;
  background-image: linear-gradient(135deg, #8bc6ec 0%, #9599e2 100%);
  border-bottom: 5px solid #b1c8f3;
}

.logo {
  margin-top: 5px;
  text-transform: uppercase;
  color: #ffffff;
  font-size: 12px;
  font-weight: 700;
}

.navbar__img {
  width: 70px;
  height: 60px;
}

@media (max-width: 768px) {
  .navbar__btns {
    display: none;
  }
}
</style>
