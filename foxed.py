import telebot
bot = telebot.TeleBot("TOKEN")
@bot.message_handler(commands=['start','help'])
def send_flex(message):
	bot.reply_to(message, "Flex")