<template>
  <!-- <pce-promo /> -->
  <section class="promo" id="about">
    <div class="promo__title">
      <img class="promo__img" src="@/assets/mumu.png" alt="#" />
      <p class="promo__title_text">
        <b>PestControlExpert</b> эксклюзивно предоставляет услуги по проведению
        дезинфекционных мероприятий: мониторинг, дезинфекции, дезинсекции,
        дератизации сети <b>кафу МУ-МУ</b>. Помимо этого, на постоянной основе -
        по договорам, помогаем более <b>50 клиентам</b> и более
        <b>100 клиентам</b> по разовым услугам на территории г. Москвы и МО.
      </p>
    </div>
  </section>

  <section class="services" id="services">
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
          Комплекс мероприятий направленных на уничтожение патогенных
          (болезнетворных) и условно-патогенных микроорганизмов в окружающей
          человека среде.
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
          Дератизационные мероприятия - включают в себя комплекс
          организационных, профилактических, истребительных мер, проводимых
          юридическими и физическими лицами, с целью ликвидации или снижения
          численности грызунов и уменьшения их вредного воздействия на человека
          и окружающую его среду.
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
          Комплекс мероприятий, направленный на уничтожение возбудителей
          инфекционных заболеваний и разрушение токсинов на объектах внешней
          среды для предотвращения попадания их на кожу, слизистые и раневую
          поверхность.
        </div>
      </div>
      <div class="services__item">
        <img src="@/assets/images/ceo.png" alt="ceo" class="services__icon" />
        <div class="services__subtitle">
          Санитарно-эпидемиологическое обследование
        </div>
        <div class="services__descr">
          Деятельность, направленная на установление соответствия
          (несоответствия) требованиям технических регламентов, государственных
          санитарно-эпидемиологических правил и нормативов производственных,
          общественных помещений, зданий, сооружений, оборудования, транспорта,
          технологического оборудования, технологических процессов, рабочих
          мест.
        </div>
      </div>
    </div>
  </section>
  <pce-price></pce-price>

  <!-- <pce-dialog-form></pce-dialog-form> -->
  <section class="callback" id="map">
    <h1>Контакты</h1>
    <div class="callback__wrapper">
      <div class="callback__form">
        <form class="form" id="form" @submit.prevent>
          <h4>Заказать обратный звонок</h4>
          <div class="callback__name">
            <div class="callback__input">
              <input
                :value="name"
                @input="name = $event.target.value"
                required
                class="input"
                type="text"
                id="name"
                placeholder="Ваше имя"
              />
            </div>
            <div class="callback__input">
              <input
                :value="phone"
                @input="phone = $event.target.value"
                required
                class="input"
                type="text"
                id="phone"
                placeholder="Ваш номер телефона"
              />
            </div>
          </div>
          <div class="callback__input">
            <textarea
              :value="comment"
              @input="comment = $event.target.value"
              required
              class="input"
              type="text"
              id="comment"
              rows="10"
              cols="35"
              placeholder="Ваш комментарий"
            >
            </textarea>
          </div>
          <div class="callback__triggers">
            <pce-button @click="sendForm" style="background: black"
              >Отправить
            </pce-button>
            <div class="callback__policy">
              <input
                :checked="checked"
                @change="checked = $event.target.checked"
                required
                type="checkbox"
              />
              <span>
                Я согласен(а) с
                <a target="/privacy" href="/privacy">
                  политикой конфидициальности
                </a>
              </span>
            </div>
          </div>
        </form>
      </div>
      <div class="callback__map">
        <h4>Карта покрытия</h4>
        <iframe
          class="callback__map_yandex"
          src="https://yandex.ru/map-widget/v1/?um=constructor%3Ae3928b383dca618308743ed387767940121ffab42b25f44f04dd31219e4cf45a&amp;source=constructor"
          width="320"
          height="348"
          frameborder="0"
        ></iframe>
      </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios';
export default {
  name: 'MainPage',
  data() {
    return {
      name: '',
      phone: '',
      comment: '',
      checked: null,
    };
  },
  methods: {
    sendForm() {
      if (this.checked === true) {
        const order = {
          id: Date.now(),
          name: this.name,
          phone: this.phone,
          comment: this.comment,
        };
        axios.post('/api/order', order);
        this.name = '';
        this.phone = '';
        this.comment = '';
        this.checked = null;
      }
    },
  },
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
  padding: 40px 20px;
  font-size: 30px;
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
  width: 230px;
  height: 230px;
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

.promo__title_text {
  padding: 20px 50px;
  font-size: 30px;
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

.callback__wrapper {
  display: flex;
  flex-wrap: wrap;
  padding-top: 15px;
}

.callback__form {
  display: flex;
}
.callback__name {
  flex-wrap: wrap;
}
.callback__map_yandex {
  border-radius: 10px;
}
@media (min-width: 1024px) {
  .callback__map_yandex {
    border: 5px solid black;
    width: 1000px;
  }
}

.form {
  display: flex;
  flex-direction: column;
}
.callback__input {
  display: flex;
  flex-direction: column;
  padding: 10px;
}

.input {
  padding: 10px;
  border-radius: 10px;
}
.callback__triggers {
  flex-direction: column-reverse;
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 20px;
}
.callback__name {
  display: flex;
}
</style>
