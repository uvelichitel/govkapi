package govkapi

import (
	"net/http"
	"time"
)

//Общие параметры
//Каждый метод API имеет собственный набор входных параметров. Помимо этого, существуют общие параметры, которые могут быть использованы во всех методах:
//
//    lang — определяет язык, на котором будут возвращаться различные данные, например, названия стран и городов. Может содержать строковое обозначение языка или его идентификатор (получить идентификатор языка пользователя Вы можете методом account.getInfo).
//        ru (0) — русский,
//        ua (1) — украинский,
//        be (2) — белорусский,
//        en (3) — английский,
//        es (4) — испанский,
//        fi (5) — финский,
//        de (6) — немецкий,
//        it (7) — итальянский.
//    Кириллические имена будут автоматически транслитерированы в латиницу для всех языков, кроме русского, украинского и белорусского.
//    https — передайте 1, чтобы получать ссылки на фотографии и другие медиафайлы с протоколом https .
//    test_mode=1 — тестовый режим, позволяет выполнять запросы из нативного приложения без его включения на всех пользователей.

//Коды ошибок
//1	Произошла неизвестная ошибка.
//Попробуйте повторить запрос позже.
//2	Приложение выключено.
//Необходимо включить приложение в настройках https://vk.com/editapp?id={Ваш API_ID} или использовать тестовый режим (test_mode=1)
//3	Передан неизвестный метод.
//Проверьте, правильно ли указано название вызываемого метода: http://vk.com/dev/methods.
//4	Неверная подпись.
//Проверьте правильность формирования подписи запроса: https://vk.com/dev/api_nohttps
//5	Авторизация пользователя не удалась.
//Убедитесь, что Вы используете верную схему авторизации. Для работы с методами без префикса secure Вам нужно авторизовать пользователя одним из этих способов: http://vk.com/dev/auth_sites, http://vk.com/dev/auth_mobile.
//6	Слишком много запросов в секунду.
//Задайте больший интервал между вызовами или используйте метод execute. Подробнее об ограничениях на частоту вызовов см. на странице http://vk.com/dev/api_requests.
//7	Нет прав для выполнения этого действия.
//Проверьте, получены ли нужные права доступа при авторизации. Это можно сделать с помощью метода account.getAppPermissions.
//8	Неверный запрос.
//Проверьте синтаксис запроса и список используемых параметров (его можно найти на странице с описанием метода).
//9	Слишком много однотипных действий.
//Нужно сократить число однотипных обращений. Для более эффективной работы Вы можете использовать execute или JSONP.
//10	Произошла внутренняя ошибка сервера.
//Попробуйте повторить запрос позже.
//11	В тестовом режиме приложение должно быть выключено или пользователь должен быть залогинен.
//Выключите приложение в настройках https://vk.com/editapp?id={Ваш API_ID}
//14	Требуется ввод кода с картинки (Captcha).
//Процесс обработки этой ошибки подробно описан на отдельной странице.
//15	Доступ запрещён.
//Убедитесь, что Вы используете верные идентификаторы, и доступ к контенту для текущего пользователя есть в полной версии сайта.
//16	Требуется выполнение запросов по протоколу HTTPS, т.к. пользователь включил настройку, требующую работу через безопасное соединение.
//Чтобы избежать появления такой ошибки, в Standalone-приложении Вы можете предварительно проверять состояние этой настройки у пользователя методом account.getInfo.
//17	Требуется валидация пользователя.
//Действие требует подтверждения — необходимо перенаправить пользователя на служебную страницу для валидации.
//18	Страница удалена или заблокирована.
//Страница пользователя была удалена или заблокирована
//20	Данное действие запрещено для не Standalone приложений.
//Если ошибка возникает несмотря на то, что Ваше приложение имеет тип Standalone, убедитесь, что при авторизации Вы используете redirect_uri=https://oauth.vk.com/blank.html. Подробнее см. http://vk.com/dev/auth_mobile.
//21	Данное действие разрешено только для Standalone и Open API приложений.
//23	Метод был выключен.
//Все актуальные методы ВК API, которые доступны в настоящий момент, перечислены здесь: http://vk.com/dev/methods.
//24	Требуется подтверждение со стороны пользователя.
//100	Один из необходимых параметров был не передан или неверен.
//Проверьте список требуемых параметров и их формат на странице с описанием метода.
//101	Неверный API ID приложения.
//Найдите приложение в списке администрируемых на странице http://vk.com/apps?act=settings и укажите в запросе верный API_ID (идентификатор приложения).
//113	Неверный идентификатор пользователя.
//Убедитесь, что Вы используете верный идентификатор. Получить ID по короткому имени можно методом utils.resolveScreenName.
//150	Неверный timestamp.
//Получить актуальное значение Вы можете методом utils.getServerTime.
//200	Доступ к альбому запрещён.
//Убедитесь, что Вы используете верные идентификаторы (для пользователей owner_id положительный, для сообществ — отрицательный), и доступ к запрашиваемому контенту для текущего пользователя есть в полной версии сайта.
//201	Доступ к аудио запрещён.
//Убедитесь, что Вы используете верные идентификаторы (для пользователей owner_id положительный, для сообществ — отрицательный), и доступ к запрашиваемому контенту для текущего пользователя есть в полной версии сайта.
//203	Доступ к группе запрещён.
//Убедитесь, что текущий пользователь является участником или руководителем сообщества (для закрытых и частных групп и встреч).
//300	Альбом переполнен.
//Перед продолжением работы нужно удалить лишние объекты из альбома или использовать другой альбом.
//500	Действие запрещено. Вы должны включить переводы голосов в настройках приложения.
//Проверьте настройки приложения: https://vk.com/editapp?id={Ваш API_ID}&section=payments
//600	Нет прав на выполнение данных операций с рекламным кабинетом.
//603	Произошла ошибка при работе с рекламным кабинетом.

const DefaultRedirectUri string = `https://oauth.vk.com/blank.html`
const OauthUri string = `https://oauth.vk.com/access_token`

type AuthParameters struct {
	ClientID     string
	ClientSecret string
	RedirectUri  string
	Display      string
	Scope        string
	ResponceType string
	APIVersion   string
	State        string
	Revoke       string
	AuthCode     string
}

var SeviceAccessToken string

type VKClient struct {
	*http.Client
	ClientID     string
	ClientSecret string
	Scope        string
	APIVersion   string
	State        string
	Revoke       string
	AuthCode     string
	AccessToken
}

type AccessToken struct {
	Token     string
	ExpiresIn time.Time
}

func MewVKClient() VKClient {

}

func (cl VKClient) Authorise() {

}

func (cl VKClient) Do(dst interface{}, Method string, MethodName string, Parameters ...string) {

}

//Implicit Flow

//Authorization Code Flow