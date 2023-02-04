import { storyblokInit, apiPlugin, getStoryblokApi } from "@storyblok/react";

storyblokInit({
  accessToken: "IfpyonLoMjLkLq8B0xu9HAtt",
  // bridge: false,
  use: [apiPlugin],
  // components: {
  //   page: Page,
  //   teaser: Teaser,
  //   grid: Grid,
  //   feature: Feature,
  // },
});

const storyBlokClient = getStoryblokApi();
window.storyBlokClient = storyBlokClient

const resultParser = (array, lat, lon) => {
  const distanceSorter = (a, b)=>{
    if(a.distance > b.distance) return 1
    if(a.distance < b.distance) return -1
    return 0
  }
  array = array.map(d=>({...d.content, distance : Math.sqrt((d.content.latitude - lat)**2 + (d.content.longitude - lon)**2)}))
  let a1 = array.filter(e=>e.is_open==1)
  let a2 = array.filter(e=>e.is_open==0)
  a1.sort(distanceSorter)
  a2.sort(distanceSorter)
  console.log('parsed', a1.concat(a2))
  return a1.concat(a2)
}

const sbQueryAroundUser = async (query)=>{
  const lat = parseFloat(localStorage.getItem("latitude"))
  const lon = parseFloat(localStorage.getItem("longitude"))
  const res = await storyBlokClient.get('cdn/stories', {
      filter_query: {
              latitude: {"lt_float": lat+0.01, "gt_float": lat-0.01},
              longitude: {"lt_float": lon+0.01, "gt_float": lon-0.01},
              __or : [
                  {stall_name: {like: `*${query}*`}},
                  {offering: {like: `*${query}*`}}
              ]
                  
          }
  })
  return resultParser(res['data']['stories'], lat, lon)
}

const getStore = async (storeId) => {
  const res = await storyBlokClient.get(`cdn/stories/${storeId}`)
  return res?.data?.story?.content
}
window.sbQueryAroundUser = sbQueryAroundUser
window.getStore = getStore
// testing
// (async ()=>{
//   console.log('getting stories')
//   const data = await storyBlokClient.get('cdn/stories')
//   console.log('from ', data)
// })()

export {
  storyBlokClient,
  sbQueryAroundUser,
  getStore
}
