/* eslint-disable */
<template>
  <div>
    <b-container fluid>
      <b-card border-variant="">
        <template slot="header">
          <b-row>
            <b-col sm="5">
              <h4 class="card-title">Информация о покупателях</h4>
            </b-col>
          </b-row>
        </template>
        <b-row>
          <b-col md="6" class="my-1">
            <b-form-group horizontal class="mb-0">
              <b-input-group>
                <b-form-input v-model="filter" placeholder="Введите id покупателя для поиска" />
                <b-input-group-append>
                  <b-btn :disabled="!filter" @click="filter = ''">Очистить</b-btn>
                </b-input-group-append>
              </b-input-group>
            </b-form-group>
          </b-col>
        </b-row>
        <b-table id="users"
                 striped
                 show-empty
                 :items="items"
                 :fields="fields"
                 :filter="filter"
                 :current-page="currentPage"
                 :per-page="perPage"
                 :total-rows="totalRows"
                 :busy.sync="isBusy"
                 ref="table"
        >
          <template slot="give_recommendation" slot-scope="row">
            <b-button size="sm" @click = "isShowingDetail = true" @click.stop="submit(row)" class="mr-2">
              {{ row.detailsShowing ? 'Скрыть' : 'Получить'}} рекомендации
            </b-button>
          </template>

          <template slot="show_details" slot-scope="row">
            <b-button size="sm" @click = "isShowingDetail = false" @click.stop="row.toggleDetails" class="mr-2">
              {{ row.detailsShowing ? 'Скрыть' : 'Показать'}} историю покупок
            </b-button>
          </template>

          <template slot="row-details" slot-scope="row">
            <b-card v-if="!isShowingDetail">
              <b-table id="products"
                       striped
                       show-empty
                       :items="row.item.items_array"
                       :fields="fields_2"
                       ref="table_2"
              ></b-table>
              <b-button size="sm" @click = "isShowingDetail = false" @click.stop="row.toggleDetails">Скрыть</b-button>
            </b-card>

            <b-card v-if="isShowingDetail">
              <b-table id="recommendations"
                       striped
                       show-empty
                       :items="recommendations"
                       :fields="fields_3"
                       :current-page="currentPage_2"
                       :per-page="perPage_2"
                       :total-rows="totalRows_2"
                       ref="table_3"
              ></b-table>
              <b-row>
                <b-col sm="12">
                  <b-pagination align="right" :total-rows="totalRows_2" :per-page="perPage_2" v-model="currentPage_2"/>
                </b-col>
              </b-row>
              <b-button size="sm" @click = "isShowingDetail = true" @click.stop="row.toggleDetails">Скрыть</b-button>
            </b-card>
          </template>
        </b-table>
        <b-row>
          <b-col sm="12">
            <b-pagination align="right" :total-rows="totalRows" :per-page="perPage" v-model="currentPage"/>
          </b-col>
        </b-row>
      </b-card>
    </b-container>
  </div>
</template>

<script>
  export default {
    name: 'Users',
    data () {
      return {
        formUrl: 'http://localhost:5001',
        fields: [
          {
            label: 'Покупатель',
            key: 'visitorid_string',
          },
          {
            label: 'Рекомендации',
            key: 'give_recommendation',
          },
          {
            label: 'История покупок',
            key: 'show_details',
          },
        ],
        fields_2: [
          {
            label:'Продукт',
            key:'itemid_string',
          },
          {
            label:'Количество покупок',
            key:'itemid_count'
          },
        ],
        fields_3: [
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
        isBusy:false,
        totalRows:1,
        currentPage:1,
        perPage:15,
        totalRows_2:1,
        currentPage_2:1,
        perPage_2:15,
        fileProducts:null,
        items:[],
        recommendations: [],
        isShowingDetail:false,
        filter: null,
      }
    },
    created() {
      this.getItems()
    },
    methods:{
      submit(row) {
        if (row.detailsShowing){
          row.toggleDetails();
          return
        }

        let myVisitor = new String(row.item.visitorid_string)
        let url = "http://localhost:5001/users/"+myVisitor;
        this.$http.get(url).then(result => {
          console.log(result);
          if (result.status === 200 || result.status === 304 ){
            if(result.body.length > 0) {
              this.recommendations = result.body;
              this.totalRows_2 = this.recommendations.length
              row.toggleDetails();
              return result.body
            }
          }

          this.recommendations = [];
          row.toggleDetails();
          return []
        },error =>{
          console.log("ERROR",error);
        });
      },

      getItems(ctx){
        let url = "http://localhost:5001/users";
        this.isBusy = true;
        return this.$http.get(url).then(result => {
          console.log(result);

          if (result.status === 200 || result.status === 304 ){
            if(result.body.length > 0) {
              this.items = result.body;
              this.totalRows = this.items.length;
              this.isBusy = false;
              return result.body
            }
          }
          this.isBusy = false;
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
