/* eslint-disable */
<template>
  <div>
    <b-container fluid>
      <b-card border-variant="">
        <template slot="header">
          <b-row>
            <b-col sm="5">
              <h4 class="card-title">Таблица истории покупок</h4>
            </b-col>
            <b-col sm="7">
              <div class ="float-right">
                <b-btn  variant="info" @click="showModalImport">загрузить файл.csv</b-btn>
              </div>
            </b-col>
          </b-row>
        </template>

        <b-row>
          <b-col md="6" class="my-1">
            <b-form-group horizontal label="Фильтр" class="mb-0">
              <b-input-group>
                <b-form-input v-model="filter" placeholder="Введите информацию для поиска" />
                <b-input-group-append>
                  <b-btn :disabled="!filter" @click="filter = ''">Очистить</b-btn>
                </b-input-group-append>
              </b-input-group>
            </b-form-group>
          </b-col>
        </b-row>
        <b-table id="home"
                 striped
                 show-empty
                 :items="items"
                 :fields="fields"
                 :current-page="currentPage"
                 :per-page="perPage"
                 :total-rows="totalRows"
                 :busy.sync="isBusy"
                 :filter="filter"
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
    <b-modal id="modal-import"
             ref="modalImport"
             title="import"
             @ok="importProducts"
             ok-title = "upload"
             cancel-title = "cancel"
             centered>
      <b-form-file v-model="fileProducts" class="mt-1"></b-form-file>
    </b-modal>
  </div>
</template>

<script>
  export default {
    name: 'Home',
    data () {
      return {
        formUrl: 'http://localhost:5001',
        fields: [
          {
            key: 'timestamp',
          },
          {
            key: 'visitorid',
          },
          {
            key: 'event',
          },
          {
            key: 'itemid',
          },
          {
            key: 'transactionid',
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
        fileProducts:null,
        items:[],
        filter: null,
      }
    },
    created() {
      this.getItems()
    },
    methods:{
      getItems(ctx){
        let url = "http://localhost:5001/events";
        this.isBusy = true;
        return this.$http.get(url).then(result => {
          Console.log(result);

          if (result.status === 200 || result.status === 304 ){
            if(result.body.length > 0) {
              this.items = result.body;
              this.isBusy = false
              this.totalRows = this.items.length;
              return result.body
            }
          }
          this.isBusy = false
          return []
        },error =>{
          this.isBusy = false;
          Console.log("ERROR",error);
        });
      },
      delete: function (url, data, callback) {
        return this.$http.delete(url,data,null).then(result => {
          callback(result.body);
        },error =>{
          Console.log("ERROR",error);
          if (error.status === 422){
            callback(error.body);
          }
          return
        });
      },
      post: function (url, data, callback) {
        Console.log(data)
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
      importProducts(){
        let formData = new FormData();
        formData.append('file', this.fileProducts);
        let url = this.formUrl+"/import";
        this.$http.post(url, formData, null).then(result => {
          Console.log(result);
          if (result.status === 200) {
            this.$refs.table.refresh();
          }
        },error =>{
          Console.log(error);
        });
      },
      showModalImport () {
        this.$refs.modalImport.show()
      },

    },
  }

</script>
