<template>
  <v-app>
    <v-app-bar
      app
      color="indigo"
      dark
    >
      <div>
        <a href="https://mineuniverse.net" style="text-decoration: none; color:white">Mineuniverse Systems</a>
      </div>

      <v-spacer></v-spacer>

    </v-app-bar>

    <v-main>
      <v-container
        class="fill-height"
        fluid
      >
        <router-view />
      </v-container>
    </v-main>

    <v-footer app>
      <span>&copy; {{ new Date().getFullYear() }}</span>
      <v-spacer></v-spacer>
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'App',

  data: () => ({
    //
  }),

  created () {
    this.$vuetify.theme.dark = true
  }
})
</script>
