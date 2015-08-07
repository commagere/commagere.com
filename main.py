import os
from datetime import date
import jinja2
import webapp2

JINJA_ENVIRONMENT = jinja2.Environment(
    loader=jinja2.FileSystemLoader(os.path.dirname(__file__)),
    extensions=['jinja2.ext.autoescape'],
    autoescape=True)

class MainHandler(webapp2.RequestHandler):
    def get(self):
        
        current_year = date.today().year
        
        template_values = {
            'current_year': current_year,
        }

        template = JINJA_ENVIRONMENT.get_template('views/index.html')
        self.response.write(template.render(template_values))
    
app = webapp2.WSGIApplication([
    ('/', MainHandler)
], debug=False)
