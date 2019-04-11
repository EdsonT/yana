<template>
  <v-container grid-list-md>
    <v-form ref="form">
      <v-text-field 
        :counter="10" 
        :rules="titleRules"
        required 
        id="title" 
        label="Post Title:" 
        v-model="title"
        ></v-text-field>
      <v-layout wrap justify-space-between>
        <v-flex xs12>
          <v-text-field required id="title" label="Post Title:" v-model="title"></v-text-field>
          <v-textarea id="description" outline label="Description:" v-model="description"></v-textarea>
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
    
      </v-layout>

      <v-btn @click="onSubmit">Submit</v-btn>
      <v-btn color="warning">Cancel</v-btn>
    </v-form>
    
  </v-container>
</template>
<script>
import axios from "axios";
export default {
  data: () => ({
    items: ["Remote", "Local"],
    title: "",
    description: "",
    company: "",
    location: "",
    type: "asd",
    titleRules:[
      v=>!!v ||'Title is Required',
      v=>v.length <=10 || 'Title must be less than 10 characters'
    ]

  }),
  methods: {
    onSubmit() {
      axios
        .post("http://localhost:8080/posts/create", {
          title: this.title,
          description: this.description,
          company: this.company,
          location: this.location,
          type: this.type,
          status: "asd",
          errors: new Errors()
        })
        .then(response => alert(response))
        .catch(error => this.errors=error.response.data);
    }
  }
  
};
</script>
<style lang="">
</style>