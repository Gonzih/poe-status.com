(ns poe.core
  (:require [reagent.core :as r]))

(enable-console-print!)

(def state (r/atom nil))

(defn root-component []
  [:div.section
   [:div.container
    [:div.column
     [:table.table.is-bordered
      [:tbody
       (for [item @state]
       [:tr
        [:td (str item)]])]]]]])

(defn init! []
  (r/render [root-component]
            (.getElementById js/document "root")))

(init!)

(->
  (js/fetch "/data.json")
  (.then (fn [response]
           (if (.-ok response)
             (.json response)
             (throw (js/Error. "Something went wrong")))))
  (.then (fn [json]
           (reset! state (js->clj json))
           (js/console.log json))))
