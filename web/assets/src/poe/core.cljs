(ns poe.core
  (:require [clojure.walk :refer [keywordize-keys]]
            [clojure.string :as strings]
            [reagent.core :as r]))

(enable-console-print!)

(def flags
  {"amsterdam" "🇳🇱"
   "australia" "🇦🇺"
   "brazil" "🇧🇷"
   "california" "🇺🇸"
   "frankfurt" "🇩🇪"
   "japan" "🇯🇵"
   "london" "🇬🇧"
   "milan" "🇮🇹"
   "moscow" "🇷🇺"
   "russia" "🇷🇺"
   "paris" "🇫🇷"
   "singapore" "🇸🇬"
   "texas" "🇺🇸"
   "washington d.c." "🇺🇸"})

(defn navbar-component []
  [:nav.navbar {:role "navigation" :aria-label "main navigation"}
   [:div.container
    [:div.navbar-brand
     [:a.navbar-burger.burger {:role "button" :arial-label "menu" :aria-expanded "false" :data-target "navbarBurger"}]
     [:div#navbarBurger.navbar-menu
      [:div.navbar-start
       [:a.navbar-item {:on-click #(prn "about")} "About"]]]]]])

(defonce state (r/atom nil))
(defonce active-tab (r/atom "PC"))

(defn server-status [{:keys [up_evidence down_evidence]}]
  (let [up-noe (or (:number_of_samples up_evidence) 0)
        down-noe (or (:number_of_samples down_evidence) 0)
        up? (>= up-noe down-noe)]
    (if up?
      (if (< up-noe (* down-noe 3))
        :unstable
        :up)
      :down)))

(defn server-state-component [{:keys [up_evidence down_evidence] :as server}]
  (let [up-noe (or (:number_of_samples up_evidence) 0)
        down-noe (or (:number_of_samples down_evidence) 0)
        status (server-status server)
        title (str up-noe " positive reports, " down-noe " negative reports")]
    [:td
     (case status
       :up       [:span.tag.is-success {:title title} "Server is up"]
       :down     [:span.tag.is-danger {:title title} "Server is down"]
       :unstable [:span.tag.is-warning {:title title} "Server is unstable"])]))

(defn server-component [{:keys [server_name platform] :as server}]
  [:tr
   [:td
    [:span.flag (flags (strings/lower-case server_name))]
    server_name]
   [server-state-component server]])

(defn server-table-component []
  [:table.table.is-stripped
   [:tbody
    (let [data (filter #(= (:platform %) @active-tab)
                       @state)]
      (doall
        (for [server data]
          ^{:key (str (:server_name server)
                      (:platform server))}
          [server-component server])))]])

(defn servers-component []
  [:div.container
   [:div.tabs
    [:ul
     (doall
       (for [fil ["PC" "XBOX"]]
         ^{:key fil}
         [:li {:class (when (= @active-tab fil) "is-active")}
          [:a {:on-click #(reset! active-tab fil)} fil " login servers"]]))]]
   [server-table-component]])

(defn content-component []
  [:div.content
   [:h2 "What is this site?"]
   [:hr]
   [:p
    "We continuously monitor the status of "
    [:a {:href "https://www.pathofexile.com/" :target "_blank"} "Path of Exile"]
    " servers. If there are any interruptions in service, it should be visible on this page."]])

(defn notification-component []
  (let [len (count @state)
        statuses (map server-status @state)
        num-down (count (filter #(= :down %) statuses))
        num-unstable (count (filter #(= :unstable %) statuses))
        overall-status (if (zero? num-down)
                          (if (zero? num-unstable)
                            :up
                            :unstable)
                          :down)
        klasses {:up "is-success" :down "is-danger" :unstable "is-warning"}
        msgs {:up "All servers are operational" :down "Some servers might be down" :unstable "Some servers might be unstable"}
        icons {:up "fa-check-circle" :down "fa-times-circle" :unstable "fa-exclamation-circle"}]
    [:div.notification {:class (overall-status klasses)}
     [:i.fas {:class (overall-status icons)}]
     (overall-status msgs)]))

(defn footer-component []
  [:footer.footer
   [:div.content.has-text-centered
    [:span
     [:a {:href "http://poe-status.com" }"poe-status.com"]
     " by "
     [:a {:href "https://github.com/Gonzih" :target "_blank"} "Max Gonzih"]]]
   [:div.content.has-text-centered.is-small
    [:span
     "Powered by "
     [:a {:href "https://golang.org/" :target "_blank"} "Go"]
     ", "
     [:a {:href "https://clojurescript.org/" :target "_blank"} "ClojureScript"]
     ", "
     [:a {:href "https://bulma.io/" :target "_blank"} "Bulma"]
     "."]
    [:br]
    [:span "This site is fan-made and not affiliated with Grinding Gear Games in any way."]]])

(defn root-component []
  [:div
   [:div.container
    [:div.section
     ; [navbar-component]
     [notification-component]
     [content-component]
     [servers-component]]]
   [footer-component]])

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
    (js/setInterval load-data! 20000)
    (load-data!)))
