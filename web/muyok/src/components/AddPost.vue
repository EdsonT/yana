<template>
  <v-container grid-list-md>
    <v-layout wrap justify-space-between>
      <v-flex xs12>
        <v-form ref="form">
          <v-text-field id="title" label="Post Title:" v-model="title"></v-text-field>
          <v-textarea id="description" outline label="Description:" v-model="description"></v-textarea>
        </v-form>
      </v-flex>
      <v-flex xs12>
        <v-text-field id="company" label="Company" v-model="company"></v-text-field>
      </v-flex>
      <v-flex xs6>
        <v-text-field id="location" label="Location" v-model="location"></v-text-field>
      </v-flex>
      <v-flex xs6>
        <v-select id="type" :items="items" label="Type" v-model="type"></v-select>
      </v-flex>
      <v-flex>
        <v-btn color="success" @click="onSubmit">Create</v-btn>
      </v-flex>
      <v-flex>
        <v-btn color="warning">Cancel</v-btn>
      </v-flex>
    </v-layout>
  </v-container>
</template>
<script>
export default {
  data: () => ({
    items: ["Remote", "Local"],
    title: "",
    description: "",
    company: "",
    location: "",
    type: ""
  }),
  methods: {
    onSubmit(){
      alert('submitting')
    }
  },
};

</script>
<style lang="">
</style>