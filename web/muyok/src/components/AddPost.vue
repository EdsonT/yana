<template>
  <v-container grid-list-xl>
    <v-layout wrap justify-space-between>
      <v-flex xs12>
        <v-form ref="form">
          <v-text-field v-model="model" label="Post Title:"></v-text-field>
          <v-textarea outline label="Description:"></v-textarea>
        </v-form>
      </v-flex>

      <v-flex xs12 md8>
        <v-slider v-model="max" label="Max characters"></v-slider>

        <v-checkbox v-model="allowSpaces" label="Allow spaces"></v-checkbox>

        <v-text-field v-model="match" label="Value must match"></v-text-field>
      </v-flex>
      <v-flex xs12 md4>
        <v-text-field v-model="model" :counter="max" :rules="rules" label="First name"></v-text-field>
      </v-flex>
    </v-layout>
  </v-container>
</template>
<script>
export default {};
</script>
<style lang="">
</style>