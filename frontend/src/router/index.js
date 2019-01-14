import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Products from '@/components/Products'
import Users from '@/components/Users'
import Recommendations from '@/components/Recommendations'
//import NeuralNetwork from '@/components/NeuralNetwork'

Vue.use(Router)

const routes = [
    {
      name: 'Home',
      path: '/',
      component: Home,
    },
    {
      name: 'Users',
      path: '/users',
      component: Users,
    },
    {
      name: 'Products',
      path: '/products',
      component: Products,
    },
    {
      name: 'Recommendations',
      path: '/recommendations',
      component: Recommendations,
    },
  //  {
  //    name: 'NeuralNetwork',
  //    path: '/neuralnetwork',
  //    component: NeuralNetwork,
  //  },
    ];

const router = new Router({routes,mode:'history'});

export default router;
