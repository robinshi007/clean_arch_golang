<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated class="">
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          @click="leftDrawerOpen = !leftDrawerOpen"
          aria-label="Menu"
          icon="menu"
        />
        <q-toolbar-title>
          Quasar App
        </q-toolbar-title>

        <q-space />
          <q-btn flat label="Login" :to="{name: 'auth.login'}" v-if="!isLoggedIn" />
        <q-btn-dropdown stretch flat :label="name" v-if="isLoggedIn">
          <q-list>
            <q-item tabindex="0">
              <q-item-section>
                <q-item-label>{{email}}</q-item-label>
              </q-item-section>
            </q-item>
            <q-item tabindex="1" clickable @click="onLogout">
              <q-item-section>
                <q-item-label>Logout</q-item-label>
              </q-item-section>
            </q-item>
          </q-list>
        </q-btn-dropdown>

        <!-- <div>Quasar v{{ $q.version }}</div> -->
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
      content-class="bg-grey-2"
    >
    <q-list class="app-menu">
        <q-item-label header></q-item-label>
      <q-item clickable :to="{ name: 'home.index' }" exact>
          <q-item-section>
            <q-item-label>Home</q-item-label>
            <q-item-label caption>Home Page</q-item-label>
          </q-item-section>
        </q-item>
        <q-item clickable :to="{ name: 'home.about' }">
          <q-item-section>
            <q-item-label>About</q-item-label>
            <q-item-label caption>About Page</q-item-label>
          </q-item-section>
        </q-item>
        <q-expansion-item clickable label="Admin">
          <q-item clickable :to="{ name: 'admin.account.list' }">
            <q-item-section>
              <q-item-label>Account</q-item-label>
            </q-item-section>
          </q-item>
        </q-expansion-item>
      </q-list>
    </q-drawer>

    <q-page-container>
      <transition
      name="fade"
      mode="out-in"
      :duration="250"
      @leave="resetScroll"
      >
        <router-view />
      </transition>
    </q-page-container>
  </q-layout>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  name: 'LayoutDefault',
  data() {
    return {
      leftDrawerOpen: false,
    };
  },
  computed: {
    ...mapGetters([
      'isLoggedIn',
      'email',
      'name',
    ]),
  },
  methods: {
    onLogout() {
      this.$store.dispatch('logout');
      // reload page after logout
      window.location.reload();
      this.$q.notify({ message: 'Logout successfully.' });
    },
    resetScroll(el, done) {
      document.documentElement.scrollTop = 0;
      document.body.scrollTop = 0;
      done();
    },
  },
};
</script>

<style>
.fade-enter-active,
.fade-leave-active {
  transition-duration: 0.25s;
  transition-property: opacity;
  transition-timing-function: ease;
}

.fade-enter,
.fade-leave-to {
  opacity: 0
}
.q-notification {
  margin: 4px 10px 0;
}
.app-menu .q-expansion-item--expanded > div > .q-item > .q-item__section--main {
  color: $primary;
  font-weight: 700;
}
.app-menu .q-expansion-item__content .q-item {
  border-radius: 0 10px 10px 0;
  margin-right: 12px;
}
.app-menu .q-item.q-router-link--active {
  background: #e3f2ff;
}
</style>
