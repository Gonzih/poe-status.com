(ns poe.core
  (:require [reagent.core :as r]
            [clojure.walk :refer [keywordize-keys]]))

(enable-console-print!)

(def state (r/atom nil))

(defn server-state [{:keys [up_evidence down_evidence]}]
  (let [up-noe (or (:number_of_samples up_evidence) 0)
        down-noe (or (:number_of_samples down_evidence) 0)
        up? (>= up-noe down-noe)
        label (if up? "Up" "Down")
        klass (if up? "fa-check" "fa-exclamation-triangle")
        title (str up-noe " positive reports, " down-noe " negative reports")]
    [:td
     (if up?
       (if (pos? down-noe)
         [:span.tag.is-warning {:title title} "Server is unstable"]
         [:span.tag.is-success {:title title} "Server is up"])
       [:span.tag.is-danger {:title title} "Server is down"])]))

(defn server-component [{:keys [server_name platform] :as server}]
  [:tr
   [:td platform]
   [:td server_name]
   [server-state server]])

(defn root-component []
  [:div.section
   [:div.container
    [:div.column
     [:table.table.is-stripped
      [:tbody
       (for [server @state]
         ^{:key (str (:server_name server)
                     (:platform server))}
         [server-component server])]]]]])

(defn init! []
  (r/render [root-component]
            (.getElementById js/document "root")))

(init!)


(defn load-data! []
  (->
    (js/fetch "/data.json")
    (.then (fn [response]
             (if (.-ok response)
               (.json response)
               (throw (js/Error. "Something went wrong")))))
    (.then (fn [json]
             (reset! state (keywordize-keys (js->clj json)))
             (js/console.log json)))))

; (js/setInterval load-data! 60000)
(load-data!)
