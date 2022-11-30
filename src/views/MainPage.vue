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

  <section class="advantages">
    <h2 class="advantages__title">НАШИ УСЛУГИ</h2>
    <div class="advantages__wrapper">
      <div class="advantages__item">
        <img
          src="@/assets/images/cockroach.png"
          alt="cup"
          class="advantages__icon"
        />
        <div class="advantages__subtitle">Название услуги</div>
        <div class="advantages__descr">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Vero velit
          illo maiores ipsam. Sint reiciendis enim non blanditiis totam illo
          quia unde molestiae officia, voluptate autem. Consectetur ullam ex
          molestias.
        </div>
      </div>
      <div class="advantages__item">
        <img
          src="@/assets/images/rat.png"
          alt="heart"
          class="advantages__icon advantages__icon_animated"
        />
        <div class="advantages__subtitle">Название услуги</div>
        <div class="advantages__descr">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Vero velit
          illo maiores ipsam. Sint reiciendis enim non blanditiis totam illo
          quia unde molestiae officia, voluptate autem. Consectetur ullam ex
          molestias.
        </div>
      </div>
      <div class="advantages__item">
        <img
          src="@/assets/images/virus.png"
          alt="shoe"
          class="advantages__icon"
        />
        <div class="advantages__subtitle">Название услуги</div>
        <div class="advantages__descr">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Vero velit
          illo maiores ipsam. Sint reiciendis enim non blanditiis totam illo
          quia unde molestiae officia, voluptate autem. Consectetur ullam ex
          molestias.
        </div>
      </div>
    </div>
  </section>

  <div v-show="timer" class="callnowbutton">
    <a class="phone" href="tel:+7999999999">
      <!-- <img class="phone__icon" src="@/assets/phone.svg" alt="phone" /> -->
      Связаться с нами
    </a>
  </div>
</template>

<script>
export default {
  name: 'MainPage',
  data() {
    return {
      timer: false,
    };
  },
  methods: {
    changeVisible() {
      setTimeout(() => {
        this.timer = true;
      }, 1000);
    },
  },
  mounted() {
    this.changeVisible();
  },
};
</script>

<style>
.promo {
  min-height: 500px;
  width: 100%;
  padding: 20px 50px 20px 50px;
  overflow: hidden;
}

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

.callnowbutton {
  transition: all 0.5s all 0.5s;
  transform: rotate(90deg);
  position: fixed;
  z-index: 100;
  padding: 10px;
  border-radius: 20px 20px 0 0;
  left: -45px;
  bottom: 80px;
  height: 30px;
  text-align: center;
  font-weight: 700;
  font-size: 10px;
  overflow: hidden;
  background: #00bb00;
  display: flex;
  justify-content: center;
  align-items: center;
}
.phone {
  text-decoration: none;
  color: #fff;
}
.advantages {
  padding: 0 0 80px 0;
  /* background-color: #f6f6f6; */
}
.advantages__wrapper {
  margin-top: 65px;
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}
.advantages__item {
  display: flex;
  flex-direction: column;
  align-items: center;
  max-width: 238px;
  margin-top: 10px;
}
.advantages__title {
  color: #fff;
}
.advantages__subtitle {
  margin-top: 38px;
  color: #ffffff;
  font-size: 18px;
  font-weight: 700;
}
.advantages__descr {
  margin-top: 26px;
  color: #ffffff;
  font-size: 14px;
  font-weight: 300;
  text-align: center;
}
.advantages__descr a {
  color: #f00;
  text-decoration: underline;
}

.advantages__icon {
  width: 100%;
  height: 100%;
}

.advantages__icon_animated {
  animation-name: heartbeat;
  animation-duration: 2s;
  animation-timing-function: ease;
  animation-iteration-count: infinite;
}
.advantages__icon_animated:hover {
  animation-play-state: paused;
}
@keyframes heartbeat {
  from {
    transform: scale(1);
  }
  50% {
    transform: scale(1.1);
  }
  to {
    transform: scale(1);
  }
}
</style>
