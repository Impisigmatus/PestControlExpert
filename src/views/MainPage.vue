<template>
  <!-- <pce-promo /> -->
  <section class="promo">
    <div class="promo__title">
      <img class="promo__img" src="@/assets/mumu.png" alt="#" />
      <p>
        <b>Pest Control Expert</b> эксклюзивно предоставляет услуги сети кафе
        Му-му. Помимо этого помогли более <b>100 клиентам</b> по Москве и
        Московской области
      </p>
    </div>
  </section>

  <section class="services">
    <h2 class="services__title">НАШИ УСЛУГИ</h2>
    <div class="services__wrapper">
      <div class="services__item">
        <img
          src="@/assets/images/cockroach.png"
          alt="cockroach"
          class="services__icon"
        />
        <div class="services__subtitle">Дезинсекция</div>
        <div class="services__descr">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Vero velit
          illo maiores ipsam. Sint reiciendis enim non blanditiis totam illo
          quia unde molestiae officia, voluptate autem. Consectetur ullam ex
          molestias.
        </div>
      </div>
      <div class="services__item">
        <img
          src="@/assets/images/rat.png"
          alt="rat"
          class="services__icon services__icon_animated"
        />
        <div class="services__subtitle">Дератизация</div>
        <div class="services__descr">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Vero velit
          illo maiores ipsam. Sint reiciendis enim non blanditiis totam illo
          quia unde molestiae officia, voluptate autem. Consectetur ullam ex
          molestias.
        </div>
      </div>
      <div class="services__item">
        <img
          src="@/assets/images/virus.png"
          alt="virus"
          class="services__icon"
        />
        <div class="services__subtitle">Дезинфекция</div>
        <div class="services__descr">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Vero velit
          illo maiores ipsam. Sint reiciendis enim non blanditiis totam illo
          quia unde molestiae officia, voluptate autem. Consectetur ullam ex
          molestias.
        </div>
      </div>
      <div class="services__item">
        <img src="@/assets/images/ceo.png" alt="ceo" class="services__icon" />
        <div class="services__subtitle">СЭО</div>
        <div class="services__descr">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Vero velit
          illo maiores ipsam. Sint reiciendis enim non blanditiis totam illo
          quia unde molestiae officia, voluptate autem. Consectetur ullam ex
          molestias.
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'MainPage',
};
</script>

<style>
.promo {
  min-height: 500px;
  width: 100%;
  padding: 20px 50px 20px 50px;
  overflow: hidden;
  background-color: black;
  background: url(@/assets/images/services_bg.jpg) center/cover no-repeat;
}

/* @media (max-width: 321px) {
  .promo {
    background: url(@/assets/bg.jpg);
  }
} */

.promo__img {
  width: 200px;
  height: 200px;
}
.promo__title {
  color: white;
}

.phone__icon {
  width: 50px;
  height: 50px;
}

.services {
  padding: 0 0 80px 0;
  background: radial-gradient(
    circle at 24.1% 68.8%,
    rgb(50, 50, 50) 0%,
    rgb(0, 0, 0) 99.4%
  );
}
.services__wrapper {
  margin-top: 25px;
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}
.services__item {
  display: flex;
  flex-direction: column;
  align-items: center;
  max-width: 238px;
  margin-top: 10px;
}
.services__title {
  padding: 15px;
  color: #fff;
}
.services__subtitle {
  margin-top: 38px;
  color: #ffffff;
  font-size: 18px;
  font-weight: 700;
}
.services__descr {
  margin-top: 26px;
  color: #ffffff;
  font-size: 14px;
  font-weight: 300;
  text-align: center;
}
.services__descr a {
  color: #f00;
  text-decoration: underline;
}

.services__icon {
  width: 100%;
  height: 100%;
}

.services__icon_animated {
  animation-name: heartbeat;
  animation-duration: 2s;
  animation-timing-function: ease;
  animation-iteration-count: infinite;
}
.services__icon_animated:hover {
  animation-play-state: paused;
}

@keyframes heartbeat {
  from {
    transform: rotateX(1);
  }
  50% {
    transform: rotateX(1.1);
  }
  to {
    transform: rotateX(1);
  }
}
</style>
