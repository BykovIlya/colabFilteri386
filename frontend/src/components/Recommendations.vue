/* eslint-disable */
<template>
  <div>
    <b-container fluid>
      <b-card border-variant="">
        <template slot="header">
          <b-row>
            <b-col sm="5">
              <h4 class="card-title">Персональные рекомендации для покупателя {{nameOfRecommendation}}</h4>
            </b-col>
            <div class ="float-left">
              <b-btn v-b-modal.modalPrevent>Введите id покупателя</b-btn>
            </div>
          </b-row>
        </template>
        <b-table id="recommendations"
                 striped
                 show-empty
                 :items="items"
                 :fields="fields"
                 :current-page="currentPage"
                 :per-page="perPage"
                 :total-rows="totalRows"
                 :busy.sync="isBusy"
                 ref="table"
        >
        </b-table>
        <b-row>
          <b-col sm="12">
            <b-pagination align="right" :total-rows="totalRows" :per-page="perPage" v-model="currentPage"/>
          </b-col>
        </b-row>
      </b-card>
    </b-container>
    <b-modal id="modalPrevent"
             ref="modal"
             title="Enter the user"
             @ok="handleOk"
             @shown="clearName">
      <form @submit.stop.prevent="handleSubmit">
        <b-form-input type="text"
                      placeholder="Enter the user"
                      v-model="nameOfRecommendation"></b-form-input>
      </form>
    </b-modal>
  </div>
</template>

<script>
  export default {
    name: 'Recommendations',
    data () {
      return {
        formUrl: 'http://localhost:5001',
        fields: [
          {
            label: 'Продукт',
            key: 'product',
          },
          {
            label: 'Рейтинг рекомендации',
            key: 'mpRating',
          },
        ],
        newItem: {
          id:null,
          name:'',
        },
        isBusy: false,
        totalRows:1,
        currentPage:1,
        perPage:15,
        fileProducts:null,
        nameOfRecommendation: '',
        items:[]
      }
    },
    created() {
      this.getItems()
    },
    methods:{
      clearName () {
        this.nameOfRecommendation = ''
      },
      handleOk (evt) {
        evt.preventDefault()
        if (!this.nameOfRecommendation) {
          alert('Please enter the user')
        } else {
          this.handleSubmit()
        }
      },
      handleSubmit () {
        let myVisitor = new String(this.nameOfRecommendation)
        let url = "http://localhost:5001/recommendations/"+myVisitor;
        this.$http.get(url).then(result => {
            console.log(result);
            this.$refs.modal.hide()
            this.getItems()
        })
      },
      getItems(ctx){
        if (this.nameOfRecommendation.length ==0){
          this.items = [];
          return;
        }

        let url = "http://localhost:5001/recommendations";
        this.isBusy = true;
        return this.$http.get(url).then(result => {
          console.log(result);

          if (result.status === 200 || result.status === 304 ){
            if(result.body.length > 0) {
              this.items = result.body;
              this.totalRows = this.items.length;
              this.isBusy = false
              return result.body
            }
          }
          thi.isBusy = false
          this.items = [];
          return []
        },error =>{
          this.isBusy = false;
          console.log("ERROR",error);
        });
      },
      delete: function (url, data, callback) {
        return this.$http.delete(url,data,null).then(result => {
          callback(result.body);
        },error =>{
          console.log("ERROR",error);
          if (error.status === 422){
            callback(error.body);
          }
          return
        });
      },
      post: function (url, data, callback) {
        console.log(data)
        return this.$http.post(url,data,null).then(result => {
          callback(result);
        },error =>{
          callback(error);
          return
        });
      },
      put: function (url, data, callback) {
        return this.$http.put(url,data,null).then(result => {
          callback(result);
        },error =>{
          callback(error);
          return
        });
      },

    },
  }

</script>
