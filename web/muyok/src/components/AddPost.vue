<template>
  <form >
    <v-text-field
      v-model="title"
      v-validate="'required|max:10'"
      :counter="10"
      :error-messages="errors.collect('title')"
      label="Title"
      data-vv-name="title"
      required
    ></v-text-field>
    <span v-text="errorsResponse.get('Title')"></span>
    <v-text-field
      v-model="company"
      v-validate="'required|max:10'"
      :counter="10"
      :error-messages="errors.collect('company')"
      label="Company Name"
      data-vv-name="company"
      required
    ></v-text-field>
    <v-text-field
      v-model="location"
      v-validate="'required|max:10'"
      :counter="10"
      :error-messages="errors.collect('location')"
      label="Location"
      data-vv-name="location"
      required
    ></v-text-field>
      <v-select
      v-model="type"
      v-validate="'required'"
      :items="items"
      :error-messages="errors.collect('type')"
      label="Job Type"
      data-vv-name="type"
      required
    ></v-select>

    <v-btn @click="submit">submit</v-btn>
    <v-btn @click="clear">clear</v-btn>
  </form>
</template>

<script>
import axios from "axios";
  class ErrorsResponse{
    constructor(){
      this.errorsResponse={ };
    }
    get(field){
      if(this.errorsResponse[field]){
         return this.errorsResponse[field];
      }
    }
    record(errorsResponse){
      this.errorsResponse=errorsResponse;
    }
  }
  export default {
    $_veeValidate: {
      validator: 'new'
    },

    data: () => ({
      title:'',
      company:'',
      location: '',
      status:'active',
      type:null,
      items:[
        'Remote Job',
        'Local Job'
      ],
      dictionary: {
        attributes: {
          email: 'E-mail Address'
          // custom attributes
        },
        custom: {
          title: {
            required: () =>this.errorsResponse.get('Title'),
            max: 'The Title field may not be greater than 10 characters'
            // custom messages
          },
          select: {
            required: 'Select field is required'
          }
        }
      },
      errorsResponse:new ErrorsResponse()
    }),

    mounted () {
      this.$validator.localize('en', this.dictionary)
      
    },

    methods: {
      submit () {
        this.$validator.validateAll()
        axios
        .post("http://localhost:8080/posts/create", {
          title: this.title,
          description: this.description,
          company: this.company,
          location: this.location,
          type: this.type,
          status: this.status
        }).then(response => alert(response))
          .catch(error=>this.errorsResponse.record(error.response.data.InnerErrors));
        
      },
      clear () {
        this.title='',
        this.company='',
        this.location= "",
        this.status='',
        this.type=null,
      
        this.$validator.reset()
      }
    }
  }
</script>