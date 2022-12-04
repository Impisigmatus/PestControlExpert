<template>
  <div class="navbar">
    <div class="navbar__logo">
      <img
        class="navbar__img"
        src="@/assets/logo.png"
        alt="#"
        @click="$router.push('/')"
      />
      <div class="navbar__title">
        <h1>Pest Control Expert</h1>
      </div>
    </div>

    <div class="navbar__contact">
      <a class="navbar__phone" href="tel:+79057247681">+7 (905) 724-76-81</a>
      <a class="navbar__email" href="email:pestcontrolexpert@yandex.ru"
        >pestcontrolexpert@yandex.ru
      </a>
    </div>
    <div class="navbar__btns">
      <pce-button class="navbar__btn"
        ><a class="navbar__link" href="#form"
          >Заказать обратный звонок
        </a></pce-button
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
  flex-direction: row;
  flex-wrap: wrap;
  align-items: center;
  padding: 0 15px;
  background: radial-gradient(
    circle at 24.1% 68.8%,
    rgb(50, 50, 50) 0%,
    rgb(0, 0, 0) 99.4%
  );
  border-bottom: 3px solid #ffffff;
}

.navbar__title {
  margin-top: 5px;
  text-transform: uppercase;
  color: #ffffff;
  font-size: 12px;
  font-weight: 700;
}

.navbar__contact {
  display: flex;
  flex-direction: column;
}
.navbar__phone,
.navbar__email {
  color: white;
}
.navbar__logo {
  display: flex;
  align-items: center;
}
.navbar__link {
  color: white;
}
.navbar__link:hover {
  color: black;
}

.navbar__btn {
  color: black;
}

.navbar__btn:hover {
  color: black;
}
.navbar__img {
  width: 70px;
  height: 60px;
}

@media (max-width: 768px) {
  .navbar {
    height: 150px;
  }
  .navbar__btns {
    display: none;
  }
  .navbar__logo {
    display: flex;
    align-items: center;
    flex-direction: column;
  }
  .navbar__img {
    width: 70px;
    height: 60px;
  }
  .navbar__title {
    font-size: 10px;
  }
}
</style>
