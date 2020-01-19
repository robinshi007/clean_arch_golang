<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated class="fixed-top">
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
      content-class="bg-write"
    >
    <q-scroll-area style="height: calc(100% - 50px); margin-top: 50px">
      <q-list class="app-menu">
        <!-- <q-item-label header></q-item-label> -->
        <q-item clickable :to="{ name: 'home.index' }" exact>
          <q-item-section>
            <q-item-label>Home</q-item-label>
          </q-item-section>
        </q-item>
        <q-item clickable :to="{ name: 'home.about' }">
          <q-item-section>
            <q-item-label>About</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />
        <q-expansion-item
          clickable
          label="Admin"
          active-class="active"
          default-opened
          >
          <q-item dense clickable :to="{ name: 'admin.account.list' }">
            <q-item-section>
              <q-item-label>Account</q-item-label>
            </q-item-section>
          </q-item>
          <q-item dense clickable :to="{ name: 'admin.user.list' }">
            <q-item-section>
              <q-item-label>User</q-item-label>
            </q-item-section>
          </q-item>
          <q-item dense clickable :to="{ name: 'admin.redirect.list' }">
            <q-item-section>
              <q-item-label>Redirect</q-item-label>
            </q-item-section>
          </q-item>
        </q-expansion-item>
      </q-list>
    </q-scroll-area>

      <div class="absolute-top bg-white layout-drawer-toolbar">
        <form autocorrect="off" autocapitalize="off" autocomplete="off" spellcheck="false">
          <q-input class=" full-width doc-algolia bg-primary"
            ref="docAlgolia"
            dense
            square
            dark
            borderless
            placeholder="Search"
            >
            <template v-slot:append>
              <q-icon name="search" />
            </template>
          </q-input>
        </form>
        <div class="layout-drawer-toolbar__shadow absolute-full no-pointer-events"></div>
      </div>
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
    ...mapGetters({
      isLoggedIn: 'auth/isLoggedIn',
      email: 'auth/email',
      name: 'auth/name',
    }),
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
body {
  background-color: #fafafa;
}
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
  color: #027be3;
  font-weight: 700;
}
.app-menu .q-expansion-item__content .q-item {
  border-radius: 0 10px 10px 0;
  margin-right: 12px;
}
.app-menu .q-item.q-router-link--active {
  background: #e3f2ff;
}
.q-expansion-item--expanded>div>.q-item {
  color: #027be3;
}
.absolute-top, .absolute-full, .fixed-top {
  top: 0;
  left: 0;
  right: 0;
}
.absolute-top  {
  position: absolute;
}
.fixed-top {
  position: fixed;
}
.doc-algolia {
  padding: 0 18px 0 16px;
  height: 50px;
}
.layout-drawer-toolbar {
  margin-right: -2px;
}
.doc-algolia .q-icon {
  color: #fafafa;
}
.layout-drawer-toolbar__shadow{
  bottom:-10px
}
.layout-drawer-toolbar__shadow:after{
  content:"";
  position:absolute;
  top:0;
  right:0;
  bottom:10px;
  left:0;
  box-shadow:0 0 10px 2px rgba(0,0,0,0.2),0 0px 10px rgba(0,0,0,0.24);
}
.no-pointer-events {
  pointer-events: none!important;
}
</style>
