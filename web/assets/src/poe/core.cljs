(ns poe.core
  (:require [reagent.core :as r]
            [clojure.walk :refer [keywordize-keys]]))

(enable-console-print!)

(defn navbar []
  [:nav.navbar {:role "navigation" :aria-label "main navigation"}
   [:div.container
    [:div.navbar-brand
     [:a.navbar-burger.burger {:role "button" :arial-label "menu" :aria-expanded "false" :data-target "navbarBurger"}]
     [:div#navbarBurger.navbar-menu
      [:div.navbar-start
       [:a.navbar-item {:on-click #(prn "about")} "About"]]]]]])

(defonce state (r/atom nil))
(defonce active-tab (r/atom "PC"))

(defn server-state [{:keys [up_evidence down_evidence]}]
  (let [up-noe (or (:number_of_samples up_evidence) 0)
        down-noe (or (:number_of_samples down_evidence) 0)
        up? (>= up-noe down-noe)
        label (if up? "Up" "Down")
        klass (if up? "fa-check" "fa-exclamation-triangle")
        title (str up-noe " positive reports, " down-noe " negative reports")]
    [:td
     (if up?
       (if (< up-noe (* down-noe 3))
         [:span.tag.is-warning {:title title} "Server is unstable"]
         [:span.tag.is-success {:title title} "Server is up"])
       [:span.tag.is-danger {:title title} "Server is down"])]))

(defn server-component [{:keys [server_name platform] :as server}]
  [:tr
   [:td server_name]
   [server-state server]])

(defn server-table []
  [:table.table.is-stripped
   [:tbody
    (let [data (filter #(= (:platform %) @active-tab)
                       @state)]
      (doall
        (for [server data]
          ^{:key (str (:server_name server)
                      (:platform server))}
          [server-component server])))]])

(defn root-component []
  [:div
   [navbar]
   [:div.section
    [:div.container
     [:div.column
      [:div.tabs
       [:ul
        (doall
          (for [fil ["PC" "XBOX"]]
            ^{:key fil}
            [:li {:class (when (= @active-tab fil) "is-active")}
             [:a {:on-click #(reset! active-tab fil)} fil]]))]]
      [server-table]]]]])

(defn init! []
  (r/render [root-component]
            (.-body js/document)))

(init!)

(defn load-data! []
  (->
    (js/fetch "/data.json")
    (.then (fn [response]
             (if (.-ok response)
               (.json response)
               (throw (js/Error. "Something went wrong")))))
    (.then (fn [json]
             (reset! state (keywordize-keys (js->clj json)))))))

(defonce interval
  (do
    (js/setInterval load-data! 60000)
    (load-data!)))
